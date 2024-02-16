// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello:", req.Host)
		fmt.Fprintln(res, "Method:", req.Method)
		fmt.Fprintln(res, "Proto:", req.Proto)

		for key, val := range req.Header {
			fmt.Fprintln(res, key, val)
		}

		fmt.Fprintln(res, "Host:", req.Host)
	})

	go func() {
		log.Fatal(http.ListenAndServe(":9002", nil))
	}()

	log.Fatal(http.ListenAndServe(":9003", nil))
}
