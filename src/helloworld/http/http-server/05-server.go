// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		header := res.Header() // get response headers
		header.Set("Content-Type", "application/json")
		header.Set("User-Agent", "Mozilla/5.0")
		header.Set("sec-ch-ua-platform", "Windows 10")

		res.Header().Set("Date", "02/11/2024")
		res.Header().Add("NewARbitraryHeader", "You are awesome!")
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadrequest == 400
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

	log.Fatal(http.ListenAndServe(":9000", nil))
}

// scp 05-server.go root@194.58.102.129:/root/go-servers
