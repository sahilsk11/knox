package resolver

import (
	"encoding/json"

	"github.com/sahilsk11/knox/internal/domain/player"
)

type listDevicesResponse struct {
	Devices []player.PlayerDevice `json:"devices"`
}

func (m httpServer) listDevices([]byte) ([]byte, error) {
	devices, err := m.PlayerService.ListAvailableDevices()
	if err != nil {
		return nil, err
	}
	response := listDevicesResponse{
		Devices: devices,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}
