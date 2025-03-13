package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// new HTTP-server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello-mello, tosy-bosy, Huggy-Wuggy")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// channel for receiving end signals
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("Error shutting down server: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	fmt.Printf("Starting at port: %s\n", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	fmt.Println("Server stopped")
}
