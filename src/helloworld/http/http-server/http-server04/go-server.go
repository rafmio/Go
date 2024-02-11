// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello-mello, Tosy-Bosy, Kissy-Missy")
	})

	// handle `/hello/golang` route to `http.DefaultServeMux`
	http.HandleFunc("/hello/golang", func(res http.ResponseWriter,
		req *http.Request) {
			fmt.Fprint(res, "<b>Hello for 'hello/golang'\n</b>")
		})

	http.HandleFunc("/htmlcode", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "<h1>Try HTML code</h>")
	})

	// listen and serve using `http.DefaultServeMux`
	http.ListenAndServe(":9000", nil)
}
