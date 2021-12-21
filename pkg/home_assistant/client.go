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
	State      LightState
}

type ControlLightsResponse struct {
	EntityName  string     `json:"entity_id"`
	State       LightState `json:"state"`
	LastChanged time.Time  `json:"last_changed"`
	LastUpdated time.Time  `json:"last_updated"`
}

type LightState string

const (
	LightState_On  LightState = "ON"
	LightState_Off LightState = "OFF"
)

func (m client) ControlLights(input ControlLightsInput) error {
	requestBody := map[string]string{
		"state": strings.ToLower(string(input.State)),
	}
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", m.BaseURL+"/api/states/"+input.EntityName, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", "Bearer "+m.AccessToken)

	response, err := m.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBytes))

	return nil
}
