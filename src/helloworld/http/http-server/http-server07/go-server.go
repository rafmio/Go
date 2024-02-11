// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two goroutines to `wg` WaitGroup
	wg.Add(2)

	// create a default route handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: " + req.Host)
		fmt.Fprint(res, "\nYour HTTP version is: " + req.Proto)
	})

	// goroutine to launch a server on port 9000
	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil))
		wg.Done()	//	one goroutine finished
		}()

	// goroutine to launch a server on port 9001
	go func() {
		log.Fatal(http.ListenAndServe(":9001", nil))
		wg.Done()
		}()

		// wait until WaitGroup is Done
		wg.Wait()
}
