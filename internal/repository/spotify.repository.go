package repository

import (
	"context"
	"log"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2/clientcredentials"

	domain "github.com/sahilsk11/knox/internal/domain/player"
)

type spotifyRepository struct {
	Client *spotify.Client
}

func NewSpotifyRepository(spotifyClientID, spotifyClientSecret string) domain.PlayerVendorRepository {
	config := &clientcredentials.Config{
		ClientID:     spotifyClientID,
		ClientSecret: spotifyClientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(context.Background(), token)

	return spotifyRepository{
		Client: spotify.New(httpClient),
	}
}
