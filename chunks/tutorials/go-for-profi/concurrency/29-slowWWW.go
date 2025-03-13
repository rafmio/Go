package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	delay := random(0, 15)
	time.Sleep(time.Duration(delay) * time.Second)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Delay: %d\n", delay)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	// seed := time.Now().Unix()
	// rand.Seed(seed)

	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("No port provided, using default port %s\n", PORT)
	} else {
		PORT = ":" + arguments[1]
		fmt.Printf("Using port %s\n", PORT)
	}

	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
