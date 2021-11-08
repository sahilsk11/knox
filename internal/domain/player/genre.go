package player

import (
	"fmt"
	"math/rand"
	"time"
)

type PlaybackGenre string

var (
	PlaybackGenre_SoftPop PlaybackGenre = "SOFT_POP"
	PlaybackGenre_Rap     PlaybackGenre = "RAP"
)

var genreArtists map[PlaybackGenre][]string = map[PlaybackGenre][]string{
	PlaybackGenre_SoftPop: {
		"Glass Animals",
		"BROCKHAMPTON",
		"Kota the Friend",
	},
	PlaybackGenre_Rap: {
		"Travis Scott",
		"Drake",
		"Lil Uzi Vert",
		"J. Cole",
		"Kendrick Lamar",
		"21 Savage",
	},
}

var genreTracks map[PlaybackGenre][]string = map[PlaybackGenre][]string{
	PlaybackGenre_SoftPop: {
		"affection",
		"vampire",
		"Guilty Conscience",
		"The Difference",
	},
	PlaybackGenre_Rap: {
		"m y . l i f e",
		"Already Best Friends",
	},
}

type PlaybackGenerateSeedResponse struct {
	Artists []string
	Tracks  []string
}

func GenerateSeed(genre PlaybackGenre) PlaybackGenerateSeedResponse {
	artists := genreArtists[genre]
	shuffle(artists)
	numArtists := random(1, 5)
	selectedArtists := artists[:min(numArtists, len(artists))]

	tracks := genreTracks[genre]
	shuffle(tracks)
	numTracks := 5 - numArtists
	selectedTracks := tracks[:min(numTracks, len(tracks))]

	fmt.Println(PlaybackGenerateSeedResponse{
		Artists: selectedArtists,
		Tracks:  selectedTracks,
	})

	return PlaybackGenerateSeedResponse{
		Artists: selectedArtists,
		Tracks:  selectedTracks,
	}
}

func random(min, max int) int {
	max += 1
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func shuffle(vals []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
