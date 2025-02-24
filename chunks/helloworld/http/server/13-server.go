package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// call setHeaders()
	setHeaders(w, r)

	// call printHeaders()
	printHeaders(w, r)

	// write response body
	fmt.Fprintf(w, "Hello, World! This is a simple HTTP server written in Go.")
}

func setHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")             // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS") // methods 'PUT', 'PATCH' and 'DELETE' have been deleted
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, hx-request, hx-target, hx-current-url")
	fmt.Fprintf(w, "Headers set successfully.")
}

func printHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Headers received:\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Remote Address: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)

	http.ListenAndServe(":8080", mux) // start server on port 8080
}
