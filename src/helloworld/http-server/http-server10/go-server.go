// https://medium.com/rungo/secure-https-servers-in-go-a783008b36da
package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World")
	})

	// run server on port 9000
	log.Fatal(http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil))
}
