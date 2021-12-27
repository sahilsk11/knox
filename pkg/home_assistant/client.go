package home_assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client interface {
	ControlLights(ControlLightsInput) error
	ControlThermostat(input ControlThermostatInput) error
}

func NewClient(baseURL, accessToken string) Client {
	return client{
		HTTPClient:  *http.DefaultClient,
		BaseURL:     baseURL,
		AccessToken: accessToken,
	}
}

type client struct {
	HTTPClient  http.Client
	BaseURL     string
	AccessToken string
}

type ControlLightsInput struct {
	EntityName string
	SwitchType SwitchType
	State      ToggleState
	Brightness *int
}

type SwitchType string

const (
	SwitchType_Light  SwitchType = "light"
	SwitchType_Switch SwitchType = "switch"
)

type ControlLightsResponse struct {
	EntityName  string      `json:"entity_id"`
	State       ToggleState `json:"state"`
	LastChanged time.Time   `json:"last_changed"`
	LastUpdated time.Time   `json:"last_updated"`
}

type ToggleState string

const (
	ToggleState_On  ToggleState = "ON"
	ToggleState_Off ToggleState = "OFF"
)

func (m client) ControlLights(input ControlLightsInput) error {
	requestBody := map[string]interface{}{
		"entity_id": input.EntityName,
	}
	if input.Brightness != nil {
		requestBody["brightness"] = *input.Brightness
	}
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"%s/api/services/%s/turn_%s",
		m.BaseURL,
		input.SwitchType,
		strings.ToLower(string(input.State)),
	)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return fmt.Errorf("failed to create request - %s", err.Error())
	}
	request.Header.Add("Authorization", "Bearer "+m.AccessToken)

	response, err := m.HTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to complete request with %s - %s", string(jsonRequestBody), err.Error())
	}
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response - %s", err.Error())
	}

	return nil
}

type ControlThermostatInput struct {
	State             ToggleState
	TargetTemperature *int
	EntityName        string
}

func (m client) ControlThermostat(input ControlThermostatInput) error {
	requestBody := map[string]interface{}{
		"entity_id": input.EntityName,
	}

	var controlType string
	if input.State == ToggleState_On && input.TargetTemperature != nil {
		controlType = "set_temperature"
		requestBody["temperature"] = *input.TargetTemperature
	} else if input.State == ToggleState_On {
		controlType = "turn_on"
	} else {
		controlType = "turn_off"
	}

	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/api/services/climate/%s", m.BaseURL, controlType)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return fmt.Errorf("failed to create request - %s", err.Error())
	}
	request.Header.Add("Authorization", "Bearer "+m.AccessToken)

	response, err := m.HTTPClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to complete request with %s - %s", string(jsonRequestBody), err.Error())
	}
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response - %s", err.Error())
	}

	return nil
}
