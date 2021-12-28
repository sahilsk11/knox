package resolver

import (
	"github.com/gin-gonic/gin"
)

func (m httpServer) theater(c *gin.Context) {
	err := m.LightsApp.EnableTheaterScene()
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	c.JSON(200, gin.H{"success": "true"})
}
