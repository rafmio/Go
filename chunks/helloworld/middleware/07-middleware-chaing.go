package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello-mello from loggingMiddleware()\n"))
		log.Println("Request has been received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello-mello from authMiddleware()\n"))
		if r.Header.Get("Authorization") != "secret" {
			http.Error(w, "Nonauthorized", http.StatusUnauthorized)
		}
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello-mello from helloHandler()\n"))
}

func chainMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}

func main() {
	handler := chainMiddleware(
		http.HandlerFunc(helloHandler),
		loggingMiddleware,
		authMiddleware,
	)

	http.Handle("/", handler)
	log.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
