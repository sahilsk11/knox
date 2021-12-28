package resolver

import (
	"github.com/gin-gonic/gin"
)

func (m httpServer) goodnight(c *gin.Context) {
	err := m.LightsApp.GoodnightScene()
	if err != nil {
		returnErrorJson(err, c)
	}

	c.JSON(200, gin.H{"success": "true"})
}
