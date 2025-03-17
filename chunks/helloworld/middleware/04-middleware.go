package main

import (
	"fmt"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.URL)
		next.ServeHTTP(w, r)
	})
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Header", "Custom Value")
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello-mello from helloHandler()")
}

func chainMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

func main() {
	hello := http.HandlerFunc(helloHandler)

	mux := http.NewServeMux()
	mux.Handle("/", chainMiddleware(hello, loggingMiddleware, headerMiddleware))

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
