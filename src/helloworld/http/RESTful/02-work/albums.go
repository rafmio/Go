package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type album struct {
	ID     string `json: "id"`
	Title  string `json: "title"`
	Artist string `json: "artist"`
}

var albums = []album{
	{ID: "1", Title: "Master of Puppets", Artist: "Metallica"},
	{ID: "2", Title: "Taste of Honey", Artist: "Beatles"},
	{ID: "3", Title: "Reace sells but who's bying", Artist: "Megadeth"},
	{ID: "4", Title: "Resistance", Artist: "Alec Koff"},
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	// for _, val := range albums {
	// 	fmt.Fprintf(w, "%s: %s of '%s'\n", val.ID, val.Title, val.Artist)
	// }
	// w.WriteHeader(200)
	// fmt.Fprintf(w, "Success!\n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func postAlbums(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")
	artist := r.FormValue("artist")

	if id == "" || title == "" || artist == "" {
		http.Error(w, "Missing required fields", 400)
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", 400)
		return
	}

	albums = append(albums, album{ID: id, Title: title, Artist: artist})

	w.WriteHeader(201)
	fmt.Fprintf(w, "Album added successfull\n")
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method Not Allowed", 405)
		return
	} else {
		fmt.Println("Method determined as a 'DELETE'")
	}

	id := r.FormValue("id")
	fmt.Printf("id = %s\n", id)

	if id == "" {
		http.Error(w, "Missing requiered field", 400)
		return
	}

	for i, alb := range albums {
		if alb.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			break
		}
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "Album deleted successfully\n")
}

func main() {
	http.HandleFunc("/", getAlbums)
	http.HandleFunc("/add", postAlbums)
	http.HandleFunc("/del", deleteAlbum)

	http.ListenAndServe(":9003", nil)
}
