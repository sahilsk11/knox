package resolver

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilsk11/knox/internal/domain/player"
)

type listDevicesResponse struct {
	Devices []player.PlayerDevice `json:"devices"`
}

func (m httpServer) listDevices(c *gin.Context) {
	devices, err := m.PlayerService.ListAvailableDevices()
	if err != nil {
		returnErrorJson(err, c)
		return
	}
	response := listDevicesResponse{
		Devices: devices,
	}

	c.JSON(200, response)
}
