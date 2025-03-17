// Example provided by GigaChat
package main

import (
	"fmt"
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

// main handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello-mello, tosy-bosy, Huggy-Wuggy"))
}

func main() {
	index := http.HandlerFunc(indexHandler) // ?? насколько я понял мы приводим indexHandler к типу HandlerFunc
	fmt.Printf("Type of 'index': %T\n", index)

	http.Handle("/", loggingMiddleware(index))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
