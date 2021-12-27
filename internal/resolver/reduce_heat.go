package resolver

import (
	"encoding/json"
	"fmt"

	domain "github.com/sahilsk11/knox/internal/domain/thermostat"
	"github.com/sahilsk11/knox/internal/service"
)

type reduceHeatRequest struct {
	ThermostatName domain.ThermostatName `json:"thermostatName"`
}

func (m httpServer) reduceHeat(requestBody []byte) ([]byte, error) {
	input := reduceHeatRequest{}

	err := json.Unmarshal(requestBody, &input)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body %s - %s", string(requestBody), err.Error())
	}

	setTemperatureInput := service.SetTemperatureInput{
		ThermostatName:    input.ThermostatName,
		TargetTemperature: 60,
	}
	err = m.ThermostatService.SetTemperature(setTemperatureInput)
	if err != nil {
		return nil, err
	}

	response := map[string]string{
		"success": "true",
	}
	responseBody, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
