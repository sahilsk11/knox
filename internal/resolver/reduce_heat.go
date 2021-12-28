package resolver

import (
	"github.com/gin-gonic/gin"
	domain "github.com/sahilsk11/knox/internal/domain/thermostat"
	"github.com/sahilsk11/knox/internal/service"
)

type reduceHeatRequest struct {
	ThermostatName domain.ThermostatName `json:"thermostatName"`
}

func (m httpServer) reduceHeat(c *gin.Context) {
	var requestBody reduceHeatRequest

	err := c.ShouldBindJSON(requestBody)
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	setTemperatureInput := service.SetTemperatureInput{
		ThermostatName:    requestBody.ThermostatName,
		TargetTemperature: 60,
	}
	err = m.ThermostatService.SetTemperature(setTemperatureInput)
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	c.JSON(200, gin.H{"success": "true"})
}
