package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	SpotifyConfig SpotifyConfig `json:"spotify"`
}

type SpotifyConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func LoadConfig(filepath string) (*Config, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %s", err.Error())
	}

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
