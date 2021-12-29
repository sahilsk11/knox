package resolver

import (
	"github.com/gin-gonic/gin"
	domain "github.com/sahilsk11/knox/internal/domain/thermostat"
	"github.com/sahilsk11/knox/internal/service"
)

type setHeatRequest struct {
	ThermostatName    domain.ThermostatName `json:"thermostatName"`
	TargetTemperature int                   `json:"targetTemperature"`
}

func (m httpServer) setHeat(c *gin.Context) {
	var requestBody setHeatRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	setTemperatureInput := service.SetTemperatureInput{
		ThermostatName:    requestBody.ThermostatName,
		TargetTemperature: requestBody.TargetTemperature,
	}
	err = m.ThermostatService.SetTemperature(setTemperatureInput)
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	c.JSON(200, gin.H{"success": "true"})
}
