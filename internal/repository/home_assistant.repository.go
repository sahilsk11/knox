package repository

import (
	"fmt"

	domain "github.com/sahilsk11/knox/internal/domain/light_controller"
	"github.com/sahilsk11/knox/pkg/home_assistant"
)

type homeAssistantRepository struct {
	Client home_assistant.Client
}

func NewHomeAssistantRepository(accessToken, baseURL string) domain.LightControllerRepository {
	client := home_assistant.NewClient(baseURL, accessToken)
	return homeAssistantRepository{
		Client: client,
	}
}

func (m homeAssistantRepository) ControlLights(input domain.ControlLightsInput) error {
	controlLightsInput := home_assistant.ControlLightsInput{
		EntityName: "light.guest_bedroom",
		State:      home_assistant.LightState_Off,
	}
	err := m.Client.ControlLights(controlLightsInput)
	if err != nil {
		return fmt.Errorf("[home assistant repository] failed to control lights: %s", err.Error())
	}

	return nil
}
