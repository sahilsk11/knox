package repository

import (
	"fmt"

	lights_domain "github.com/sahilsk11/knox/internal/domain/light_controller"
	thermostat_domain "github.com/sahilsk11/knox/internal/domain/thermostat"
	"github.com/sahilsk11/knox/pkg/home_assistant"
)

type homeAssistantRepository struct {
	Client home_assistant.Client
}

func NewHomeAssistantRepository(accessToken, baseURL string) lights_domain.LightControllerRepository {
	client := home_assistant.NewClient(baseURL, accessToken)
	return homeAssistantRepository{
		Client: client,
	}
}

func (m homeAssistantRepository) ControlLights(input lights_domain.ControlLightsInput) error {
	controlLightsInput := home_assistant.ControlLightsInput{
		EntityName: input.Light.HomeAssistantEntityName,
		State:      home_assistant.ToggleState(input.State),
		Brightness: input.Brightness,
	}
	if input.Light.SwitchType == lights_domain.SwitchType_Switch {
		controlLightsInput.SwitchType = home_assistant.SwitchType_Switch
	} else {
		controlLightsInput.SwitchType = home_assistant.SwitchType_Light
	}
	err := m.Client.ControlLights(controlLightsInput)
	if err != nil {
		return fmt.Errorf("[home assistant repository] failed to control lights: %s", err.Error())
	}

	return nil
}

func (m homeAssistantRepository) SetTemperature(input thermostat_domain.SetTemperatureInput) error {
	controlThermostatInput := home_assistant.ControlThermostatInput{
		State:      home_assistant.ToggleState_On,
		EntityName: input.Thermostat.HomeAssistantEntityName,
	}
	err := m.Client.ControlThermostat(controlThermostatInput)
	if err != nil {
		return err
	}

	controlThermostatInput = home_assistant.ControlThermostatInput{
		TargetTemperature: &input.TargetTemperature,
	}
	err = m.Client.ControlThermostat(controlThermostatInput)
	if err != nil {
		return err
	}

	return nil
}

func (m homeAssistantRepository) TurnOffThermostat(entityName string) error {
	controlThermostatInput := home_assistant.ControlThermostatInput{
		State:      home_assistant.ToggleState_Off,
		EntityName: entityName,
	}
	err := m.Client.ControlThermostat(controlThermostatInput)
	if err != nil {
		return err
	}

	return nil
}
