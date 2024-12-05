package metadata_test

import (
	"fmt"
	"geniusapi/metadata"
	"geniusapi/models"
	"strings"
	"testing"
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
