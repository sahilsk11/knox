package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"

	domain "github.com/sahilsk11/knox/internal/domain/player"
	"github.com/sahilsk11/knox/internal/util"
)

type spotifyRepository struct {
	Client *spotify.Client
}

func NewSpotifyRepository(accessToken, refreshToken string, tokenExpiryUnix int64, clientID string) domain.PlayerVendorRepository {
	tm := time.Unix(tokenExpiryUnix, 0)
	if tm.Before(time.Now()) {
		// refresh token
		s, err := util.GetAndSaveAccessToken(
			refreshToken,
			clientID,
		)
		if err != nil {
			log.Fatal(err)
		}
		accessToken = s.AccessToken
		refreshToken = s.RefreshToken
	}
	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       tm,
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
	genreRecommendations, err := m.generateSeed(input.Genre)
	if err != nil {
		return err
	}
	firstTrack := genreRecommendations[0].URI
	err = m.Client.PlayOpt(context.Background(), &spotify.PlayOptions{
		DeviceID: (*spotify.ID)(&input.DeviceID),
		URIs:     []spotify.URI{firstTrack},
	})
	if err != nil {
		return fmt.Errorf("failed to start playback on device ID %s: %s", input.DeviceID, err.Error())
	}

	err = m.Client.Repeat(context.Background(), "context")
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

func (m spotifyRepository) searchTrack(trackName string) (*spotify.FullTrack, error) {
	resp, err := m.Client.Search(context.Background(), trackName, spotify.SearchType(spotify.SearchTypeTrack))
	if err != nil {
		return nil, err
	} else if resp.Tracks.Total < 1 {
		return nil, fmt.Errorf("search track result for \"%s\" returned 0 results", trackName)
	}

	return &resp.Tracks.Tracks[0], nil
}

func (m spotifyRepository) searchArtist(artistName string) (*spotify.FullArtist, error) {
	resp, err := m.Client.Search(context.Background(), artistName, spotify.SearchType(spotify.SearchTypeArtist))
	if err != nil {
		return nil, err
	} else if resp == nil || resp.Artists == nil {
		fmt.Println(resp.Artists)
		return nil, fmt.Errorf("nil result for artist search on \"%s\"", artistName)
	} else if resp.Artists.Total < 1 {
		return nil, fmt.Errorf("search artist result for \"%s\" returned 0 results", artistName)
	}

	return &resp.Artists.Artists[0], nil
}

func (m spotifyRepository) generateSeed(genre domain.PlaybackGenre) ([]spotify.SimpleTrack, error) {
	seedValues := domain.GenerateSeed(genre)
	artists := make([]spotify.ID, len(seedValues.Artists))
	tracks := make([]spotify.ID, len(seedValues.Tracks))

	for i, artistName := range seedValues.Artists {
		artist, err := m.searchArtist(artistName)
		if err != nil {
			return nil, err
		}
		artists[i] = artist.ID
	}
	for i, trackName := range seedValues.Tracks {
		track, err := m.searchTrack(trackName)
		if err != nil {
			return nil, err
		}
		tracks[i] = track.ID
	}

	resp, err := m.Client.GetRecommendations(
		context.Background(),
		spotify.Seeds{
			Artists: artists,
			Tracks:  tracks,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return resp.Tracks, nil
}
