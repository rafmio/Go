// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	wg.Add(2)

	// create a default route handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello: "+req.Host)
		fmt.Fprintln(res, "Your HTTP version is: "+req.Proto)
		for key, val := range req.Header {
			fmt.Fprintln(res, key, val)
		}
	})

	go func() {
		log.Fatal(http.ListenAndServe(":9004", nil))
		wg.Done()
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":9005", nil))
		wg.Done()
	}()

	wg.Wait()
}
