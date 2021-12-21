package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// hacky library to save and update access and refresh token

func GetAndSaveAccessToken(refreshToken, clientID string) (*spotifyRefreshResponse, error) {
	client := &http.Client{}

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("client_id", clientID)

	r, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	s := spotifyRefreshResponse{}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body - %s", err.Error())
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json %s - %s", string(body), err.Error())
	}

	return &s, writebackConfig(s)
}

type spotifyRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func writebackConfig(c spotifyRefreshResponse) error {
	config, err := LoadConfig("config/keys.json")
	if err != nil {
		return fmt.Errorf("failed to load config file - %s", err.Error())
	}
	config.SpotifyConfig.AccessToken = c.AccessToken
	config.SpotifyConfig.RefreshToken = c.RefreshToken
	config.SpotifyConfig.TokenExpiry = time.Now().Unix() + c.ExpiresIn

	conf, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal new config - %s", err.Error())
	}

	return os.WriteFile("config/keys.json", conf, 0644)
}
