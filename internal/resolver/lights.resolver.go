package resolver

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilsk11/knox/internal/domain/light_controller"
	"github.com/sahilsk11/knox/internal/service"
)

type setBrightnessRequestBody struct {
	LightName  light_controller.LightName `json:"lightName"`
	Brightness int                        `json:"brightness"`
}

func (m httpServer) setBrightness(c *gin.Context) {
	var requestBody setBrightnessRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		returnErrorJson(err, c)
		return
	}

	setBrightnessInput := service.SetBrightnessInput{
		LightName:  requestBody.LightName,
		Brightness: requestBody.Brightness,
	}
	err := m.LightService.SetBrightness(setBrightnessInput)
	if err != nil {
		returnErrorJson(err, c)
	}

	c.JSON(200, gin.H{"success": "true"})
}
