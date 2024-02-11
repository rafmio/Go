// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// create channels for safe exit
	exitSignal := make(chan interface{})

	// create server to run on port the 9000
	server := &http.Server {
		Addr: ":9000",
		Handler: nil,		// use `DefaultServeMux`
	}

	// close server after a few seconds
	time.AfterFunc(3 * time.Second, func() {
		fmt.Println("Close(): completed", server.Close())	// close `server`
		close(exitSignal)
	})

	// launch server
	err := server.ListenAndServe()
	fmt.Println("ListenAndServe():", err)

	// listen to `exitSignal` channels
	<- exitSignal
	fmt.Println("Main(): exit complete")
}
