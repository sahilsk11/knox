package service

import (
	"fmt"

	domain "github.com/sahilsk11/knox/internal/domain/light_controller"
	"github.com/sahilsk11/knox/internal/repository"
)

type LightService interface {
	SetBrightness(SetBrightnessInput) error
	TurnOff(domain.LightName) error
	TurnOn(domain.LightName) error
}

type lightService struct {
	LightRepository         domain.LightControllerRepository
	LightDatabaseRepository repository.LightDatabaseRepository
}

func NewLightService(lightRepository domain.LightControllerRepository, lightDatabaseRepository repository.LightDatabaseRepository) LightService {
	return lightService{
		LightRepository:         lightRepository,
		LightDatabaseRepository: lightDatabaseRepository,
	}
}

type SetBrightnessInput struct {
	LightName  domain.LightName
	Brightness int
}

func (m lightService) SetBrightness(input SetBrightnessInput) error {
	light, err := m.LightDatabaseRepository.GetRoom(input.LightName)
	if err != nil {
		return err
	} else if light.HomeAssistantEntityName == "" {
		return fmt.Errorf("%s missing home assistant name definition", string(input.LightName))
	}
	if input.Brightness == 0 {
		return m.TurnOff(input.LightName)
	}

	if light.SwitchType != domain.SwitchType_Adjustable {
		return fmt.Errorf("cannot adjust brightness of %s - switch type is not adjustable", string(input.LightName))
	}

	controlLightsInput := domain.ControlLightsInput{
		Light:      *light,
		Brightness: &input.Brightness,
		State:      domain.LightState_On,
	}
	err = m.LightRepository.ControlLights(controlLightsInput)
	if err != nil {
		return err
	}

	return nil
}

func (m lightService) TurnOff(LightName domain.LightName) error {
	light, err := m.LightDatabaseRepository.GetRoom(LightName)
	if err != nil {
		return err
	} else if light.HomeAssistantEntityName == "" {
		return fmt.Errorf("%s missing home assistant name definition", string(LightName))
	}

	controlLightsInput := domain.ControlLightsInput{
		Light: *light,
		State: domain.LightState_Off,
	}
	err = m.LightRepository.ControlLights(controlLightsInput)
	if err != nil {
		return err
	}

	return nil
}

func (m lightService) TurnOn(LightName domain.LightName) error {
	light, err := m.LightDatabaseRepository.GetRoom(LightName)
	if err != nil {
		return err
	} else if light.HomeAssistantEntityName == "" {
		return fmt.Errorf("%s missing home assistant name definition", string(LightName))
	}

	controlLightsInput := domain.ControlLightsInput{
		Light: *light,
		State: domain.LightState_On,
	}
	err = m.LightRepository.ControlLights(controlLightsInput)
	if err != nil {
		return err
	}

	return nil
}
