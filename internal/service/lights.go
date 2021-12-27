package service

import "github.com/sahilsk11/knox/internal/domain/light_controller"

type LightService interface {
	SetBrightness()
	TurnOff()
	TurnOn()
}

type lightService struct {
	LightRepository light_controller.LightControllerRepository
}

func NewLightService(lightRepository light_controller.LightControllerRepository) LightService {
	return lightService{
		LightRepository: lightRepository,
	}
}

func (m lightService) SetBrightness() {}

func (m lightService) TurnOff() {}

func (m lightService) TurnOn() {}
