package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/natecham/genius"
)

func main() {
	accessToken := "ycPo8Ic25mjQO6-Kcka0gUeeZDNNDmhCFh1o19YNPLGXY4W95_xXGuPwCDVjBeJA"
	client := genius.NewClient(nil, accessToken)

	// Create a slice of songs to search for
	songs := []string{"Yesterday", "Unforgiven"}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Randomly select one of the songs
	songTitle := songs[rand.Intn(len(songs))]

	fmt.Printf("Searching for: %s\n\n", songTitle)

	// Search for the selected song
	results, err := client.Search(songTitle)
	if err != nil {
		log.Fatalf("Searching error: %v", err)
	}

	lyricsURL := make(map[string]string)

	if results.Meta.Status == 200 {
		if len(results.Response.Hits) > 0 {
			for _, hit := range results.Response.Hits {
				song := hit.Result
				fmt.Printf("Song Title: %s\n", song.FullTitle)
				fmt.Printf("Artist: %s\n", song.PrimaryArtist.Name)
				fmt.Printf("URL: %s\n", song.URL)
				fmt.Println("---")
				lyricsURL[song.FullTitle] = song.URL
			}
		} else {
			fmt.Println("No results found.")
		}
	} else {
		fmt.Printf("Error: status %d, message: %s\n", results.Meta.Status, results.Meta.Message)
	}

	unforgivenLyrics, err := client.GetLyrics("https://genius.com/Metallica-the-unforgiven-lyrics")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// wirte string to file 'unforgiven.txt'
	err = writeStringToFile("unforgiven.txt", unforgivenLyrics)
	if err != nil {
		log.Fatalf("Writing lyric file error: %v", err)
	}

	fmt.Println("Lyric saved to 'unforgiven.txt'")
}

func writeStringToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
