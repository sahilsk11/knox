package main

import (
	"fmt"
	"log"

	"github.com/sahilsk11/knox/internal/app"
	"github.com/sahilsk11/knox/internal/domain/light_controller"
	"github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/repository"
	"github.com/sahilsk11/knox/internal/resolver"
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

	homeAssistantRepository := repository.NewHomeAssistantRepository(
		config.HomeAssistantConfig.AccessToken, config.HomeAssistantConfig.BaseURL,
	)
	lightDatabaseRepository, err := repository.NewLightDatabaseRepository("database/lights.json")
	if err != nil {
		log.Fatal(err)
	}
	thermostatDatabaseRepository, err := repository.NewThermostatDatabaseRepository("database/thermostat.json")
	if err != nil {
		log.Fatal(err)
	}

	playerService := service.NewPlayerService(spotifyRepository)
	lightService := service.NewLightService(homeAssistantRepository, lightDatabaseRepository)
	thermostatService := service.NewThermostatService(homeAssistantRepository, thermostatDatabaseRepository)

	lightApp := app.NewLightsApp(lightService, thermostatService)

	startServer(playerService, lightService, lightApp)

	sh(playerService)
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

func startServer(playerService service.PlayerService, lightService service.LightService, lightsApp app.LightsApp) {
	server := resolver.NewHTTPServer(playerService, lightService, lightsApp)
	fmt.Println("serving on port 8000")
	server.StartHTTPServer(8000)
}

func sh(playerService service.PlayerService) {}
