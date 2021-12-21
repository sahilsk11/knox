package main

import (
	"fmt"
	"log"

	"github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/repository"
	"github.com/sahilsk11/knox/internal/service"
	"github.com/sahilsk11/knox/internal/util"
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

	play(playerService)
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
		DeviceNameSimilarTo: "berry",
		Genre:               player.PlaybackGenre_Simp,
	}
	err := playerService.StartPlayback(input)
	if err != nil {
		log.Fatal(err)
	}
}
