package main

import (
	"fmt"
	"net/http"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

var albums = []album{
	{ID: "1", Title: "Master of Puppets", Artist: "Metallica"},
	{ID: "2", Title: "Taste of Honey", Artist: "Beatles"},
	{ID: "3", Title: "Peace sells but who's buying", Artist: "Megadeth"},
	{ID: "4", Title: "Resistance", Artist: "Alec Koff"},
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	for _, val := range albums {
		fmt.Fprintf(w, "%s: %s of '%s'\n", val.ID, val.Title, val.Artist)
	}
}

func main() {
	// mux := http.NewServeMux()
	// mux.Handle()
	http.HandleFunc("/", getAlbums)

	http.ListenAndServe(":9003", nil)
}
