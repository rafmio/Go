package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello-mello, tosy-bosy")
}

func main() {
	// Creating a handler that limits the request size to 1 MB
	handler := http.MaxBytesHandler(
		http.HandlerFunc(myHandler),
		1024*1024, // 1 MB
	)

	// register handler at the route "/"
	http.Handle("/", handler)

	// Start the server on port 8080
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
