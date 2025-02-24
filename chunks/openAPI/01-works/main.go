// https://dev.to/envitab/documenting-a-go-api-with-openapi-3-standard-a-practical-guide-jod
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:author"`
}

var books = []Book{
	{ID: 1, Title: "To Kill a Mockingbird", Author: "Harper Lee"},
	{ID: 2, Title: "1984", Author: "George Orwell"},
	{ID: 3, Title: "The Go Programming Language", Author: "Alan A. A. Donovan, Brian W. Kernighan"},
	{ID: 4, Title: "Designing Data-Intensive Applications", Author: "Martin Kleppmann"},
	{ID: 5, Title: "Code Complete", Author: "Steve McConnell"},
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "id")
	for _, book := range books {
		if strconv.Itoa(book.ID) == bookID {
			render.JSON(w, r, book)
			return
		}
	}

	render.Status(r, http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/books/{id}", GetBook)

	log.Println("Starting on : 3000")
	http.ListenAndServe(":3000", r)
}
