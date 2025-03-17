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
	fmt.Fprintln(w, "Hello with middleware!")
}

func main() {
	handler := http.HandlerFunc(helloHandler)
	http.Handle("/middleware", loggingMiddleware(handler))
	http.ListenAndServe(":8080", nil)
}
