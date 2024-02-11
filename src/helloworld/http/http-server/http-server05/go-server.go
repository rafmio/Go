// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	// handle `/` route `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		header.Set("User-Agent", "Ashrapov Browser")

		header.Set("sec-ch-ua-platform", "Ayaz OS")

		// reset date header (inline call)
		res.Header().Set("Date", "01/01/2023")

		// add new header
		res.Header().Add("NewArbitraryHeader", "You are awesome!")

		// set status header
		res.WriteHeader(http.StatusBadRequest)	//	http.StatusBadRequest == 400

		// respond with a JSON string
		fmt.Fprint(res, `{"status": "FAILURE"}`)

		for key, val := range req.Header {
			fmt.Fprint(res, "\n")
			fmt.Fprint(res, key, ":", val)
		}

		resHeaders := res.Header()

		fmt.Fprint(res, "\n\nResponse headers:\n")
		for key, val := range resHeaders {
			fmt.Fprint(res, "\n")
			fmt.Fprint(res, key, val)
		}
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9000", nil))
}
