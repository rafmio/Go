// Example provided by GigaChat
package main

import (
	"log"
	"net/http"
)

// logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// auth middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func compressingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(w, r)
	})
}

func chainMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello-mello from the main handler (indexHandler())"))
}

func main() {
	index := http.HandlerFunc(indexHandler)

	http.Handle("/", chainMiddleware(
		index,
		loggingMiddleware,
		authMiddleware,
		compressingMiddleware,
	))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
