package resolver

import (
	"encoding/json"
	"net/http"
	"strings"
)

type listGenresResponse struct {
	Genres []string `json:"genres"`
}

func (m httpServer) listGenres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	genres, err := m.PlayerService.ListGenres()
	if err != nil {
		json.NewEncoder(w).Encode(responseError{})
	}
	genresStr := make([]string, len(genres))
	for i, genre := range genres {
		genresStr[i] = strings.ReplaceAll(strings.ToLower(string(genre)), "_", " ")
	}
	response := listGenresResponse{
		Genres: genresStr,
	}
	json.NewEncoder(w).Encode(response)
}
