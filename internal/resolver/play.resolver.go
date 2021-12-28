package resolver

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/service"
)

type playRequest struct {
	DeviceID string `json:"device_id"`
	Genre    string `json:"genre"`
}

type playResponse struct {
}

func (m httpServer) play(c *gin.Context) {
	var requestBody playRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		returnErrorJson(err, c)
		return
	}

	genre := strings.ToUpper(strings.ReplaceAll(requestBody.Genre, " ", "_"))

	startPlaybackInput := service.StartPlaybackInput{
		DeviceFilter: service.DeviceFilter{
			DeviceID: &requestBody.DeviceID,
		},
		Genre: player.PlaybackGenre(genre),
	}
	err := m.PlayerService.StartPlayback(startPlaybackInput)
	if err != nil {
		returnErrorJson(fmt.Errorf("[play resolver] could not start playback %s", err.Error()), c)
	}

	c.JSON(200, gin.H{"success": "true"})
}
