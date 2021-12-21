package resolver

import (
	"encoding/json"
	"net/http"

	"github.com/sahilsk11/knox/internal/domain/player"
)

type listDevicesResponse struct {
	Devices []player.PlayerDevice `json:"devices"`
}

func (m httpServer) listDevices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	devices, err := m.PlayerService.ListAvailableDevices()
	if err != nil {
		json.NewEncoder(w).Encode(responseError{})
	}
	response := listDevicesResponse{
		Devices: devices,
	}
	json.NewEncoder(w).Encode(response)
}
