// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func createServer(name string, port int) *http.Server {
	// create `ServeMux`
	mux := http.NewServeMux()

	// create a default route handler
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello: " + name)
	})

	// create a new server
	server := http.Server {
		Addr: fmt.Sprintf(":%v", port),	// :{port}
		Handler: mux,
	}

	// return new server (pointer)
	return &server
}

func main() {
	// create a WaitGroup
	wg := new(sync.WaitGroup)

	// add two gorotines to `wg` WaitGroup
	wg.Add(2)

	// goroutine to launch a server on port 9000
	go func() {
		server := createServer("ONE", 9000)
		fmt.Println(server.ListenAndServe())
		wg.Done()
		}()

	// gorotine to launch a server on port 9001
	go func() {
		server := createServer("TWO", 9001)
		fmt.Println(server.ListenAndServe())
		wg.Done()
		}()

	// wait until WaitGroup is done
	wg.Wait()
}
