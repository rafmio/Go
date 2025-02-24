// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
// This code doesn't work! I don't know why
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello-mello, tosy-bosy\n")
	})

	mux.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hoba-nah! What's the people!\n")
	})

	http.ListenAndServe(":9000", mux)
}

// curl -X GET http://localhost:9000
