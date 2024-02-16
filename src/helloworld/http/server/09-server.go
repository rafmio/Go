// https://medium.com/rungo/running-multiple-http-servers-in-go-d15300f4e59f
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	exitSignal := make(chan interface{}) // create channels for safe exit

	// create server
	server := &http.Server{
		Addr:    ":9000",
		Handler: nil, // user `DefaultServeMux`
	}

	// close server after a few seconds
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Close(): completed", server.Close())
		close(exitSignal)
	})

	err := server.ListenAndServe() // launch server
	fmt.Println("ListenAndServe():", err)

	m := <-exitSignal // listen to `exitSignal` channel
	fmt.Printf("Type of 'm' is %T\n", m)
	fmt.Println("main(): exit complete")
}
