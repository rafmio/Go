package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	// for _, val := range albums {
	// 	fmt.Fprintf(w, "%s: %s of '%s'\n", val.ID, val.Title, val.Artist)
	// }

	// Выдать результат в JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)

	w.WriteHeader(200)
	fmt.Fprintf(w, "Success!\n")

}

func postAlbums(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	// Получаем данные из запроса
	id := r.FormValue("id")
	title := r.FormValue("title")
	artist := r.FormValue("artist")

	// Проверяем, являются ли значения непустыми
	if id == "" || title == "" || artist == "" {
		http.Error(w, "Missiing required fields", 400)
		return
	}

	// Проверяем, является ли ID числом
	_, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID format", 400)
		return
	}

	albums = append(albums, album{ID: id, Title: title, Artist: artist})

	w.WriteHeader(201) // w.WriteHeader() ф-я поле стр-ры ResponseWriter
	fmt.Fprintf(w, "Album added successfully\n")
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method Not Allowed", 405)
		return
	} else {
		fmt.Println("Method determined as 'DELETE'")
	}

	// Получаем ID альбома из запроса
	id := r.FormValue("id")
	fmt.Printf("id = %s\n", id)

	// Проверяем, является ли ID непустым
	if id == "" {
		http.Error(w, "Missing required field", 400)
		return
	}

	// Ищем альбом по ID
	for i, alb := range albums {
		// Удаляем альбом из слайса
		if alb.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			break
		}
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "Album deleted successfully\n")
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	// Получаем ID альбома из запроса
	id := r.FormValue("id")

	// Проверяем, является ли ID непустым
	if id == "" {
		http.Error(w, "Missing required field", 400)
		return
	}

	// Ищем альбом по ID
	for i, alb := range albums {
		if alb.ID == id {
			// Обновляем Title альбома
			albums[i].Title = r.FormValue("title")
			break
		}
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "Album updated successfully\n")
}

func main() {
	http.HandleFunc("/", getAlbums)
	http.HandleFunc("/add", postAlbums)
	http.HandleFunc("/del", deleteAlbum)
	http.HandleFunc("/upd", updateAlbum)
	http.ListenAndServe(":9003", nil)
}

// POST method:
// $ curl -X POST -d "id=6&title=Tejas&artist=ZZ Top" http://localhost:9003/add

// DELETE method:
// $ curl -X DELETE http://localhost:9003/del?id=4

// PUT method:
// $ curl -X PUT -d "id=2&title=Abbey Road" http://localhost:9003/upd
