package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello-mello from the handleRoor() wich was called by HandleFunc\n")
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080", nil) // start server on port 8080
}

/*
http HandleFunc

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    HandleFunc registers the handler function for the given pattern in
    DefaultServeMux. The documentation for ServeMux explains how patterns are
    matched.
*/
