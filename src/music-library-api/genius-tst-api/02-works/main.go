package main

import (
	"fmt"
	"log"

	"github.com/natecham/genius"
)

func main() {
	accessToken := "ycPo8Ic25mjQO6-Kcka0gUeeZDNNDmhCFh1o19YNPLGXY4W95_xXGuPwCDVjBeJA"
	client := genius.NewClient(nil, accessToken)

	// artist := "The Beatles"
	songTitle := "yesterday"

	// (c *Client) Search(q string) returns (*GeniusResponse, error)
	results, err := client.Search(songTitle)
	if err != nil {
		log.Printf("Searching error: %v", err)
	} else {
		fmt.Println(results.Meta.Status, results.Meta.Message)
	}

	fmt.Println(results.Response)
}

// type GeniusResponse struct {
// 	Meta     *Meta     `json:"meta"`
// 	Response *Response `json:"response"`
// }

// type Meta struct {
// 	Status  int    `json:"status"`
// 	Message string `json:"message"`
// }

// type Response struct {
// 	Artist      *Artist       `json:"artist"`
// 	Album       *Album        `json:"album"`
// 	AlbumTracks []*AlbumTrack `json:"tracks"`
// 	Albums      []*Album      `json:"albums"`
// 	Song        *Song         `json:"song"`
// 	Songs       []*Song       `json:"songs"`
// 	Annotation  *Annotation   `json:"annotation"`
// 	User        *User         `json:"user"`
// 	NextPage    int           `json:"next_page"`
// 	Hits        []*Hit        `json:"hits"`
// 	WebPage     *WebPage      `json:"web_page"`
// 	Sections    []Sections    `json:"sections"`
// }
