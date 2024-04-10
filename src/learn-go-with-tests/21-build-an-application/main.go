// https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server
package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}

// The HandlerFunc type is an adapter to allow the use of ordinary
// functions as HTTP handlers. If f is a function with the appropriate
// signature, HandlerFunc(f) is a Handler that calls f.
