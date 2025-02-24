package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title string
	Users []string
}

func main() {
	data := ViewData{
		Title: "User List",
		Users: []string{"Alice", "Bob", "Charlie", "Valera", "Sam", "Bob", "Thomas"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening")
	http.ListenAndServe(":8181", nil)
}
