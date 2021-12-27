package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/sahilsk11/knox/internal/domain/light_controller"
	"github.com/sahilsk11/knox/internal/service"
)

type setBrightnessRequestBody struct {
	LightName  light_controller.LightName `json:"lightName"`
	Brightness int                        `json:"brightness"`
}

func (m httpServer) setBrightness(requestBody []byte) ([]byte, error) {
	input := setBrightnessRequestBody{}

	err := json.Unmarshal(requestBody, &input)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal body %s - %s", string(requestBody), err.Error())
	}

	setBrightnessInput := service.SetBrightnessInput{
		LightName:  input.LightName,
		Brightness: input.Brightness,
	}
	err = m.LightService.SetBrightness(setBrightnessInput)
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
