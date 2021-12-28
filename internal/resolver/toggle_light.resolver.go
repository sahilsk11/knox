package resolver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	domain "github.com/sahilsk11/knox/internal/domain/light_controller"
)

type toggleLightRequestBody struct {
	State     domain.LightState `json:"state"`
	LightName domain.LightName  `json:"lightName"`
}

func (m httpServer) toggleLight(c *gin.Context) {
	var requestBody toggleLightRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		returnErrorJson(err, c)
		return
	}

	var err error
	if requestBody.State == domain.LightState_On {
		err = m.LightService.TurnOn(requestBody.LightName)
	} else if requestBody.State == domain.LightState_Off {
		err = m.LightService.TurnOff(requestBody.LightName)
	} else {
		returnErrorJson(fmt.Errorf("invalid state %s - expected ON or OFF", requestBody.State), c)
	}
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	c.JSON(200, gin.H{"success": "true"})
}
