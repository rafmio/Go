// https://metanit.com/go/web/2.1.php
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title:   "Hello, World!",
			Message: "This is a simple web page",
		}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
