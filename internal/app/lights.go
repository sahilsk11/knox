package app

import (
	domain "github.com/sahilsk11/knox/internal/domain/light_controller"
	"github.com/sahilsk11/knox/internal/service"
)

type LightsApp interface {
	EnableTheaterScene() error
	GoodnightScene() error
	DownstairsLightsOn() error
}

type lightsApp struct {
	LightService      service.LightService
	ThermostatService service.ThermostatService
}

func NewLightsApp(lightService service.LightService, thermostatService service.ThermostatService) LightsApp {
	return lightsApp{
		LightService:      lightService,
		ThermostatService: thermostatService,
	}
}

// Theater scene dims the lights downstairs
func (m lightsApp) EnableTheaterScene() error {
	adjustBrightnessInputs := []service.SetBrightnessInput{
		{
			LightName:  domain.Light_Living,
			Brightness: 60,
		},
		{
			LightName:  domain.Light_Kitchen,
			Brightness: 10,
		},
		{
			LightName:  domain.Light_Dining,
			Brightness: 10,
		},
	}

	turnOffInputs := []domain.LightName{
		domain.Light_Kitchen_Accent,
		domain.Light_Staircase,
		domain.Light_Entry,
	}

	controllerInput := asyncControlInput{
		SetBrightnessInputs: adjustBrightnessInputs,
		TurnOffInputs:       turnOffInputs,
	}
	err := m.asyncControl(controllerInput)
	if err != nil {
		return err
	}

	return nil
}

func (m lightsApp) GoodnightScene() error {
	turnOffInputs := []domain.LightName{
		domain.Light_Entry,
		domain.Light_Living,
		domain.Light_Dining,
		domain.Light_Kitchen,
		domain.Light_Kitchen_Accent,
		domain.Light_Downstairs_Bathroom,
		domain.Light_Storage,
	}
	turnOnInputs := []domain.LightName{
		domain.Light_Porch,
		domain.Light_Staircase,
		domain.Light_Sahil_Bedroom,
	}
	controllerInput := asyncControlInput{
		TurnOnInputs:  turnOnInputs,
		TurnOffInputs: turnOffInputs,
	}
	err := m.asyncControl(controllerInput)
	if err != nil {
		return err
	}
	return nil
}

func (m lightsApp) DownstairsLightsOn() error {
	controlInput := asyncControlInput{
		TurnOnInputs: []domain.LightName{
			domain.Light_Kitchen,
			domain.Light_Kitchen_Accent,
			domain.Light_Dining,
			domain.Light_Living,
			domain.Light_Entry,
		},
	}
	err := m.asyncControl(controlInput)
	if err != nil {
		return err
	}
	return nil
}

type asyncControlInput struct {
	TurnOnInputs        []domain.LightName
	TurnOffInputs       []domain.LightName
	SetBrightnessInputs []service.SetBrightnessInput
}

func (m lightsApp) asyncControl(input asyncControlInput) error {
	errChannels := make(
		[]chan error,
		len(input.TurnOnInputs)+len(input.TurnOffInputs)+len(input.SetBrightnessInputs),
	)
	for i := range errChannels {
		errChannels[i] = make(chan error)
	}

	offset := 0

	// asynchronously set all brightnesses
	for i, brightnessInput := range input.SetBrightnessInputs {
		go func(i int, brightnessInput service.SetBrightnessInput) {
			errChannels[i] <- m.LightService.SetBrightness(brightnessInput)
		}(i, brightnessInput)
	}

	offset += len(input.SetBrightnessInputs)

	// asynchronously turn off all lights
	for i, turnOffInput := range input.TurnOffInputs {
		go func(i int, name domain.LightName) {
			errChannels[i] <- m.LightService.TurnOff(name)
		}(offset+i, turnOffInput)
	}

	offset += len(input.TurnOffInputs)

	// asynchronously turn on all lights
	for i, turnOnInput := range input.TurnOnInputs {
		go func(i int, name domain.LightName) {
			errChannels[i] <- m.LightService.TurnOn(name)
		}(offset+i, turnOnInput)
	}

	var err error
	for i := range errChannels {
		chanErr := <-errChannels[i]
		if chanErr != nil {
			err = chanErr
		}
	}
	return err
}
