package resolver

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type listGenresResponse struct {
	Genres []string `json:"genres"`
}

func (m httpServer) listGenres(c *gin.Context) {
	genres, err := m.PlayerService.ListGenres()
	if err != nil {
		returnErrorJson(err, c)
		return
	}

	genresStr := make([]string, len(genres))
	for i, genre := range genres {
		genresStr[i] = strings.ReplaceAll(strings.ToLower(string(genre)), "_", " ")
	}
	response := listGenresResponse{
		Genres: genresStr,
	}

	c.JSON(200, response)
}
