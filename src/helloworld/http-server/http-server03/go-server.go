// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a new `ServMux`
	mux := http.NewServeMux()

	// handle `/` route
	mux.HandleFunc("/", func(res http.ResponseWriter,
		req *http.Request) {
			fmt.Fprint(res, "Hello-mello")
		})

	// handle `/hello/golang` route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hoba! What's the people!")
	})

	// listen and serve using `ServeMux`
	http.ListenAndServe(":9000", mux)
}
