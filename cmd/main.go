package main

import (
	"log"

	"github.com/sahilsk11/knox/internal/repository"
	"github.com/sahilsk11/knox/internal/service"
	"github.com/sahilsk11/knox/internal/util"
)

func main() {
	config, err := util.LoadConfig("config/keys.json")
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	spotifyRepository := repository.NewSpotifyRepository(config.SpotifyConfig.AccessToken, config.SpotifyConfig.RefreshToken)
	playerService := service.NewPlayerService(spotifyRepository)

	input := service.StartPlaybackInput{
		DeviceNameSimilarTo: "sahil-mbp",
	}
	err = playerService.StartPlayback(input)
	if err != nil {
		log.Fatal(err)
	}
}
