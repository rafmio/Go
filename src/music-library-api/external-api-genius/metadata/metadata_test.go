package metadata_test

import (
	"fmt"
	"geniusapi/metadata"
	"geniusapi/models"
	"strings"
	"testing"

	"github.com/natecham/genius"
)

func TestGetSongMetadata(t *testing.T) {
	// test data
	queryParams := models.QueryParams{
		Group: "Muse",
		Song:  "Supermassive Black Hole",
	}

	// calling the function to be checked
	songDetail, err := metadata.GetSongMetadata(queryParams)

	// check results
	if err != nil {
		t.Errorf("GetSongMetadata() returned an error: %v", err)
	}

	if songDetail == nil {
		t.Error("GetSongMetadata() returned a nil songDetail")
	}

	// check songDetail's fields
	if !strings.ContainsAny(songDetail.Title, "SupermassiveBlackHole") {
		t.Errorf("Expected title to contain 'Supermassive', 'Black', or 'Hole', got '%s'", songDetail.Title)
	}

	// check if 'Artist' is empty
	if songDetail.Artist == "" {
		t.Error("Expected non-empty artist name")
	}

	// check if 'Link' is empty
	if songDetail.Link == "" {
		t.Error("Expected non-empty link")
	}

	// check if 'ReleaseDate' is empty
	if songDetail.ReleaseDate == "" {
		fmt.Println("release:", songDetail.ReleaseDate)
		t.Error("Expected non-empty release date")
	}
}

func TestChooseSong(t *testing.T) {
	tests := []struct {
		name        string
		queryParams models.QueryParams
		songs       []*genius.Hit
		want        *genius.Song
	}{
		{
			name: "Match found",
			queryParams: models.QueryParams{
				Group: "The Beatles",
				Song:  "Yesterday",
			},
			songs: []*genius.Hit{
				{
					Result: &genius.Song{
						PrimaryArtist: &genius.Artist{Name: "The Beatles"},
						Title:         "Yesterday",
					},
				},
				{
					Result: &genius.Song{
						PrimaryArtist: &genius.Artist{Name: "Another Artist"},
						Title:         "Some Song",
					},
				},
			},
			want: &genius.Song{
				PrimaryArtist: &genius.Artist{Name: "The Beatles"},
				Title:         "Yesterday",
			},
		},
		{
			name: "No match found",
			queryParams: models.QueryParams{
				Group: "The Rolling Stones",
				Song:  "Paint It Black",
			},
			songs: []*genius.Hit{
				{
					Result: &genius.Song{
						PrimaryArtist: &genius.Artist{Name: "The Beatles"},
						Title:         "Yesterday",
					},
				},
				{
					Result: &genius.Song{
						PrimaryArtist: &genius.Artist{Name: "Another Artist"},
						Title:         "Some Song",
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := metadata.ChooseSong(&tt.queryParams, tt.songs)
			t.Logf("Test case: %s", tt.name)
			t.Logf("Query params: %+v", tt.queryParams)
			t.Logf("Got: %+v", got)
			t.Logf("Want: %+v", tt.want)
			if (got == nil && tt.want != nil) || (got != nil && tt.want == nil) {
				t.Errorf("ChooseSong() = %v, want %v", got, tt.want)
			}
			if got != nil && tt.want != nil {
				if got.PrimaryArtist.Name != tt.want.PrimaryArtist.Name || got.Title != tt.want.Title {
					t.Errorf("ChooseSong() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
