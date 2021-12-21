package main

import (
	"fmt"
	"log"

	"github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/repository"
	"github.com/sahilsk11/knox/internal/service"
	"github.com/sahilsk11/knox/internal/util"
	"github.com/sahilsk11/knox/resolver"
)

func main() {
	config, err := util.LoadConfig("config/keys.json")
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}

	spotifyRepository := repository.NewSpotifyRepository(
		config.SpotifyConfig.AccessToken,
		config.SpotifyConfig.RefreshToken,
		config.SpotifyConfig.TokenExpiry,
		config.SpotifyConfig.ClientID,
	)
	playerService := service.NewPlayerService(spotifyRepository)

	startServer(playerService)
}

func listDevices(playerService service.PlayerService) {
	devices, err := playerService.ListAvailableDevices()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range devices {
		fmt.Printf("- %s : %s\n", device.DeviceName, device.DeviceID)
	}
}

func play(playerService service.PlayerService) {
	input := service.StartPlaybackInput{
		DeviceNameSimilarTo: "Sahil’s MacBook Pro",
		Genre:               player.PlaybackGenre_Rap,
	}
	err := playerService.StartPlayback(input)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer(playerService service.PlayerService) {
	server := resolver.NewHTTPServer(playerService)
	server.StartHTTPServer(8000)
}
