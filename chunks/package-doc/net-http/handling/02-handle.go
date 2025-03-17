package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello-mello, Tosy-Bosy, Huggy-Wuggy")
}

func main() {
	var handler MyHandler
	http.Handle("/", &handler)
	http.ListenAndServe(":8080", nil)
}

/*
http.Handle:

Функция, которая регистрирует обработчик http.Handler

func Handle(pattern string, handler Handler)
    Handle registers the handler for the given pattern in DefaultServeMux.
    The documentation for ServeMux explains how patterns are matched.

*/
