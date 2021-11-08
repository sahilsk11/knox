package player

import (
	"testing"
)

func Test_GenerateSeed(t *testing.T) {
	resp := GenerateSeed(PlaybackGenre_Rap)
	if len(resp.Artists) < 1 || len(resp.Artists) > 5 {
		t.Fatalf("returned %d artists", len(resp.Artists))
	}
	if len(resp.Tracks) < 1 || len(resp.Tracks) > 5 {
		t.Fatalf("returned %d tracks", len(resp.Tracks))
	}
}
