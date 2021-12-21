package resolver

import (
	"encoding/json"
	"strings"
)

type listGenresResponse struct {
	Genres []string `json:"genres"`
}

func (m httpServer) listGenres([]byte) ([]byte, error) {
	genres, err := m.PlayerService.ListGenres()
	if err != nil {
		return nil, err
	}
	genresStr := make([]string, len(genres))
	for i, genre := range genres {
		genresStr[i] = strings.ReplaceAll(strings.ToLower(string(genre)), "_", " ")
	}
	response := listGenresResponse{
		Genres: genresStr,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}
