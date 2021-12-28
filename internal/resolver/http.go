package resolver

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sahilsk11/knox/internal/app"
	"github.com/sahilsk11/knox/internal/service"
)

type responseError struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func NewHTTPServer(authToken string, playerService service.PlayerService, lightService service.LightService, lightsApp app.LightsApp, thermostatService service.ThermostatService) httpServer {
	return httpServer{
		AuthToken:         authToken,
		PlayerService:     playerService,
		LightService:      lightService,
		LightsApp:         lightsApp,
		ThermostatService: thermostatService,
	}
}

type httpServer struct {
	AuthToken         string
	PlayerService     service.PlayerService
	LightService      service.LightService
	LightsApp         app.LightsApp
	ThermostatService service.ThermostatService
}

func (m httpServer) StartHTTPServer(port int) {
	router := gin.Default()
	router.Use(m.authMiddleware())

	// music player routes
	router.GET("/player/listDevices", m.listDevices)
	router.GET("/player/listGenres", m.listGenres)
	router.POST("/player/play", m.play)

	// light controller routes
	router.POST("/lights/setBrightness", m.setBrightness)

	// scene routes
	router.POST("/scenes/goodnight", m.goodnight)

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

func (m httpServer) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header authHeader
		if err := c.ShouldBindHeader(&header); err != nil {
			returnErrorJson(errors.New("missing authorization header"), c)
		}
		authToken := strings.Replace(header.Authorization, "Bearer ", "", 1)

		if authToken != m.AuthToken {
			c.AbortWithStatusJSON(403, gin.H{"error": "invalid auth header"})
		}
	}
}
