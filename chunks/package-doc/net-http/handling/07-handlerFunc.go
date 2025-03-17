package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from the HandlerFunc()")
}

func main() {
	handler := http.HandlerFunc(helloHandler)
	http.Handle("/handlerfunc", handler)
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error running server")
	}
}
