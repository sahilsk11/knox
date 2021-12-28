package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	SpotifyConfig       SpotifyConfig        `json:"spotify"`
	HomeAssistantConfig HomeAssistantConfig  `json:"home_assistant"`
	Authentication      AuthenticationConfig `json:"authentication"`
}

type SpotifyConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenExpiry  int64  `json:"tokenTime"`
}

type HomeAssistantConfig struct {
	AccessToken string `json:"accessToken"`
	BaseURL     string `json:"baseURL"`
}

type AuthenticationConfig struct {
	AuthToken string `json:"authToken"`
}

func LoadConfig(filepath string) (*Config, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %s", err.Error())
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %s", err.Error())
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %s", err.Error())
	}

	return &config, nil
}
