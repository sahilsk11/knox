package resolver

import (
	"fmt"
	"net/http"

	"github.com/sahilsk11/knox/internal/service"
)

type responseError struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func NewHTTPServer(playerService service.PlayerService) httpServer {
	return httpServer{
		PlayerService: playerService,
	}
}

type httpServer struct {
	PlayerService service.PlayerService
}

func (m httpServer) StartHTTPServer(port int) {
	http.HandleFunc("/listDevices", m.listDevices)
	http.HandleFunc("/listGenres", m.listGenres)
	http.HandleFunc("/play", m.play)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
