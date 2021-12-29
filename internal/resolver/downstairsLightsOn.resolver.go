package resolver

import (
	"github.com/gin-gonic/gin"
)

func (m httpServer) downstairsLightsOn(c *gin.Context) {
	err := m.LightsApp.DownstairsLightsOn()
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	c.JSON(200, gin.H{"success": "true"})
}
