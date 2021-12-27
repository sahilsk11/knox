package service

import (
	"fmt"

	domain "github.com/sahilsk11/knox/internal/domain/light_controller"
	"github.com/sahilsk11/knox/internal/repository"
)

type LightService interface {
	SetBrightness(SetBrightnessInput) error
	TurnOff(domain.RoomName) error
	TurnOn(domain.RoomName) error
}

type lightService struct {
	LightRepository         domain.LightControllerRepository
	LightDatabaseRepository repository.LightDatabaseRepository
}

func NewLightService(lightRepository domain.LightControllerRepository) LightService {
	return lightService{
		LightRepository: lightRepository,
	}
}

type SetBrightnessInput struct {
	RoomName  domain.RoomName
	Intensity int
}

func (m lightService) SetBrightness(input SetBrightnessInput) error {
	room, err := m.LightDatabaseRepository.GetRoom(input.RoomName)
	if err != nil {
		return err
	}
	if room.SwitchType != domain.SwitchType_Adjustable {
		return fmt.Errorf("cannot adjust brightness of %s - switch type is not adjustable", string(input.RoomName))
	}

	controlLightsInput := domain.ControlLightsInput{
		EntityName: room.HomeAssistantEntityName,
		Intensity:  &input.Intensity,
		State:      domain.LightState_On,
	}
	err = m.LightRepository.ControlLights(controlLightsInput)
	if err != nil {
		return err
	}

	return nil
}

func (m lightService) TurnOff(roomName domain.RoomName) error {
	room, err := m.LightDatabaseRepository.GetRoom(roomName)
	if err != nil {
		return err
	}

	controlLightsInput := domain.ControlLightsInput{
		EntityName: room.HomeAssistantEntityName,
		State:      domain.LightState_Off,
	}
	err = m.LightRepository.ControlLights(controlLightsInput)
	if err != nil {
		return err
	}

	return nil
}

func (m lightService) TurnOn(roomName domain.RoomName) error {
	room, err := m.LightDatabaseRepository.GetRoom(roomName)
	if err != nil {
		return err
	}

	controlLightsInput := domain.ControlLightsInput{
		EntityName: room.HomeAssistantEntityName,
		State:      domain.LightState_On,
	}
	err = m.LightRepository.ControlLights(controlLightsInput)
	if err != nil {
		return err
	}

	return nil
}
