package resolver

import (
	"fmt"
	"net/http"

	"github.com/sahilsk11/knox/internal/service"
)

func NewHTTPServer(playerService service.PlayerService) httpServer {
	return httpServer{
		PlayerService: playerService,
	}
}

type httpServer struct {
	PlayerService service.PlayerService
}

func (m httpServer) listDevices(h http.ResponseWriter, r *http.Request) {
	devices, err := m.PlayerService.ListAvailableDevices()
	if err != nil {
		fmt.Fprint(h, "request failed")
	}
	for _, device := range devices {
		fmt.Printf("- %s : %s\n", device.DeviceName, device.DeviceID)
	}
	fmt.Fprint(h, "done")
}

func (m httpServer) StartHTTPServer(port int) {
	http.HandleFunc("/listDevices", m.listDevices)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
