// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func createServer(name string, port int) *http.Server {
	mux := http.NewServeMux() // create 'ServeMux'
	fmt.Printf("Handler for '%v' server is created. Type of mux: %T\n", name, mux)

	// create a default route handler
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hola! You are:"+name)
	})

	// create a new server
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port), // :{port}
		Handler: mux,
	}

	fmt.Printf("Server '%v' is created\n", name)

	return &server // return new server (pointer)
}

func main() {
	var wg sync.WaitGroup // create a WaitGroup. Or: wg := new(sync.WaitGroup)
	wg.Add(2)             // add two goroutines to 'wg' WaitGroup

	go func() {
		server := createServer("ONE", 9005)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	go func() {
		server := createServer("TWO", 9006)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	wg.Wait()
}
