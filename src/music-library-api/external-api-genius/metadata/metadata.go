package metadata

import (
	"errors"
	"log"

	"github.com/natecham/genius"

	// "github.com/natecham/genius/models"
	"geniusapi/models"
)

func GetSongMetadata(queryParams models.QueryParams) (*models.SongDetail, error) {
	accessToken := "ycPo8Ic25mjQO6-Kcka0gUeeZDNNDmhCFh1o19YNPLGXY4W95_xXGuPwCDVjBeJA"

	// Create a new client for the Genius API
	client := genius.NewClient(nil, accessToken)
	log.Println("the client for genius.com has been created")

	// Search for the selected song
	results, err := client.Search(queryParams.Song)
	if err != nil {
		return nil, err
	}

	// Extract the song details from the first hit in the search results
	if results.Meta.Status == 200 {
		if len(results.Response.Hits) > 0 {
			song := results.Response.Hits[0].Result
			songDetail := &models.SongDetail{
				// SeqNum:      song.ID,
				ID:    song.ID,
				Title: song.FullTitle,
				// ReleaseDate: song.ReleaseDate,
				ReleaseDate: song.ReleaseDateForDisplay,
				Artist:      song.PrimaryArtist.Name,
				Link:        song.URL,
			}
			log.Println("the requested song has been extracted")

			return songDetail, nil
		} else {
			log.Printf("error extracting the song. Status: %v", results.Meta.Status)
			return nil, nil
		}
	} else {
		return nil, err
	}
}

func GetText(client *genius.NewClient) (string, error) {
	// get lyrics
	text, err := client.GetLyrics(client.Link)
	if err != nil {
		log.Println("error getting lyrics (text):", err)
		return "", err
	}

	// check if text (lyrics) is empty
	if text == "" {
		return "", errors.New("lyrics (text) variable is empty")
	}

	return text, nil
}
