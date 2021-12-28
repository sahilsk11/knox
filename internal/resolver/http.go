package resolver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sahilsk11/knox/internal/app"
	"github.com/sahilsk11/knox/internal/service"
)

type responseError struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func NewHTTPServer(playerService service.PlayerService, lightService service.LightService, lightsApp app.LightsApp, thermostatService service.ThermostatService) httpServer {
	return httpServer{
		PlayerService:     playerService,
		LightService:      lightService,
		LightsApp:         lightsApp,
		ThermostatService: thermostatService,
		Logger:            log.Default(),
	}
}

type httpServer struct {
	PlayerService     service.PlayerService
	LightService      service.LightService
	LightsApp         app.LightsApp
	ThermostatService service.ThermostatService
	Logger            *log.Logger
}

func (m httpServer) loggingMiddleware(handler func([]byte) ([]byte, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var (
			requestBody []byte
			err         error
		)

		if r.Method == "POST" {
			requestBody, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
			if err != nil {
				json.NewEncoder(w).Encode(responseError{ErrorMessage: err.Error()})
				m.Logger.Println(err.Error())
			}
		}

		responseBody, err := handler(requestBody)
		if err != nil {
			json.NewEncoder(w).Encode(responseError{
				ErrorCode:    400,
				ErrorMessage: err.Error(),
			})
			m.Logger.Println(err.Error())
		} else {
			w.Write(responseBody)
		}
	})
}

func (m httpServer) StartHTTPServer(port int) {
	router := gin.Default()
	router.Use(authMiddleware())

	// music player routes
	router.GET("/player/listDevices", m.listDevices)
	router.GET("/player/listGenres", m.listGenres)
	router.POST("/player/play", m.play)

	// light controller routes
	router.POST("/lights/setBrightness", m.setBrightness)

	// scene routes
	router.GET("/scenes/goodnight", m.goodnight)

	// climate routes
	router.POST("/climate/reduceHeat", m.reduceHeat)

	router.Run(fmt.Sprintf(":%d", port))
}

func returnErrorJson(err error, c *gin.Context) {
	c.AbortWithStatusJSON(400, gin.H{
		"error": err.Error(),
	})
}

type authHeader struct {
	Authorization string `header:"Authorization"`
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header authHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			returnErrorJson(errors.New("missing authorization header"), c)
		}
		authToken := strings.Replace(header.Authorization, "Bearer ", "", 1)
		if authToken != "hi" {
			c.AbortWithStatusJSON(403, gin.H{"error": "invalid auth header"})
		}
	}
}
