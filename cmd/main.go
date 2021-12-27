package main

import (
	"log"

	"github.com/sahilsk11/knox/internal/domain/light_controller"
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
	// homeAssistantRepository := repository.NewHomeAssistantRepository(
	// 	config.HomeAssistantConfig.AccessToken, config.HomeAssistantConfig.BaseURL,
	// )
	playerService := service.NewPlayerService(spotifyRepository)

	startServer(playerService)
}

func lights(h light_controller.LightControllerRepository) {
	err := h.ControlLights(light_controller.ControlLightsInput{})
	if err != nil {
		log.Fatal(err)
	}
}

func play(playerService service.PlayerService) {
	input := service.StartPlaybackInput{
		DeviceFilter: service.DeviceFilter{
			DeviceNameSimilarTo: util.StrPtr("Sahil’s MacBook Pro"),
		},
		Genre: player.PlaybackGenre_Rap,
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
