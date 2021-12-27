package resolver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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
	mux := http.NewServeMux()
	mux.Handle("/listDevices", m.loggingMiddleware(m.listDevices))
	mux.Handle("/listGenres", m.loggingMiddleware(m.listGenres))
	mux.Handle("/play", m.loggingMiddleware(m.play))
	mux.Handle("/setBrightness", m.loggingMiddleware(m.setBrightness))
	mux.Handle("/goodnight", m.loggingMiddleware(m.goodnight))
	mux.Handle("/reduceHeat", m.loggingMiddleware(m.reduceHeat))
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
