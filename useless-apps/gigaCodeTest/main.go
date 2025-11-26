package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time: %v", time.Now().Format(time.RFC1123))
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count: 42")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/count", countHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
