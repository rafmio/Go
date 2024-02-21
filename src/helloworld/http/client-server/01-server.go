package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\"%s</h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\"%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for %s\n", r.Host)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	userAgents := r.Header
	for _, uAgent := range userAgents {
		fmt.Fprintf(w, "Your User-Agents are: %s\n", &uAgent)
	}
}

func main() {
	port := ":9002"
	argumets := os.Args
	if len(argumets) == 1 {
		fmt.Println("Using default port number: ", port)
	} else {
		port = ":" + argumets[1]
	}

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", defaultHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error serving http:", err.Error())
		return
	}
}
