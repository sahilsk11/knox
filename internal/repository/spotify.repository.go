package repository

import (
	"context"
	"fmt"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"

	domain "github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/util"
)

type spotifyRepository struct {
	Client *spotify.Client
}

func NewSpotifyRepository(accessToken, refreshToken string) domain.PlayerVendorRepository {
	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	httpClient := spotifyauth.New().Client(context.Background(), token)

	return spotifyRepository{
		Client: spotify.New(httpClient),
	}
}

func (m spotifyRepository) ListDevices() ([]domain.PlayerDevice, error) {
	devices, err := m.Client.PlayerDevices(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to list spotify devices: %s", err.Error())
	}

	// adapt from spotify model to domain
	playerDevices := make([]domain.PlayerDevice, len(devices))
	for i, device := range devices {
		playerDevices[i] = domain.PlayerDevice{
			DeviceName: device.Name,
			DeviceID:   device.ID.String(),
		}
	}

	return playerDevices, nil
}

func (m spotifyRepository) StartPlayback(input domain.StartPlaybackInput) error {
	err := m.Client.PlayOpt(context.Background(), &spotify.PlayOptions{
		DeviceID:        (*spotify.ID)(&input.DeviceID),
		PlaybackContext: (*spotify.URI)(util.StrPtr("spotify:album:2up3OPMp9Tb4dAKM2erWXQ")),
	})
	if err != nil {
		return fmt.Errorf("failed to start playback on device ID %s: %s", input.DeviceID, err.Error())
	}
	return nil
}

func (m spotifyRepository) User() error {
	user, err := m.Client.CurrentUser(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("You are logged in as:", user.ID)
	return nil
}
