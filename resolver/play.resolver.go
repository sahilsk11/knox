package resolver

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func (m httpServer) play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		json.NewEncoder(w).Encode(responseError{ErrorMessage: err.Error()})
	}

	request := playRequest{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		json.NewEncoder(w).Encode(responseError{ErrorMessage: err.Error()})
	}
	genre := strings.ToUpper(strings.ReplaceAll(request.Genre, " ", "_"))

	logger := log.Default()
	logger.Printf("starting %s on %s", genre, request.DeviceID)

	err = m.PlayerService.StartPlayback(service.StartPlaybackInput{
		DeviceFilter: service.DeviceFilter{
			DeviceID: &request.DeviceID,
		},
		Genre: player.PlaybackGenre(genre),
	})
	if err != nil {
		json.NewEncoder(w).Encode(responseError{ErrorMessage: err.Error()})
	}

	json.NewEncoder(w).Encode(playResponse{})
}
