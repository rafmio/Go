// https://habr.com/ru/post/475390/
package main

import (
	"html/template"
	"net/http"
)

var tlp = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tlp.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":9000", mux)
}
