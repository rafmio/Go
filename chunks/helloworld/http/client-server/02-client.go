package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	userAgent := "Super-puper mega agent/ Safari/527.36"
	req, err := http.NewRequest("GET", "http://localhost:9002/time", nil)
	if err != nil {
		log.Fatal("Error making new request:", err.Error())
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error doing request:", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body:", err.Error())
	}

	fmt.Printf("%s\n", string(body))
}
