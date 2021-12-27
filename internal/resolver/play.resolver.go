package resolver

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/service"
)

type playRequest struct {
	DeviceID string `json:"device_id"`
	Genre    string `json:"genre"`
}

type playResponse struct {
}

func (m httpServer) play(requestBody []byte) ([]byte, error) {
	request := playRequest{}
	err := json.Unmarshal(requestBody, &request)
	if err != nil {
		return nil, fmt.Errorf("[play resolver] could not decode request body %s : %s", string(requestBody), err.Error())
	}
	genre := strings.ToUpper(strings.ReplaceAll(request.Genre, " ", "_"))

	m.Logger.Printf("starting %s on %s", genre, request.DeviceID)

	err = m.PlayerService.StartPlayback(service.StartPlaybackInput{
		DeviceFilter: service.DeviceFilter{
			DeviceID: &request.DeviceID,
		},
		Genre: player.PlaybackGenre(genre),
	})
	if err != nil {
		return nil, fmt.Errorf("[play resolver] could not start playback %s", err.Error())
	}

	jsonResponse, err := json.Marshal(playResponse{})
	if err != nil {
		return nil, fmt.Errorf("[play resolver] could not marshal response struct %v : %s", playResponse{}, err.Error())
	}

	return jsonResponse, nil
}
