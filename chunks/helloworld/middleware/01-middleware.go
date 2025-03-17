package main

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Tosy-Bosy, Huggy-Wuggy")
}

func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(helloHandler)))
	http.ListenAndServe(":8080", nil)
}
