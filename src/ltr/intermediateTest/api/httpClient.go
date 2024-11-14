// simple http-client that makes an HTTP GET request to a "http://localhost:8082/logtracker/fetch_entries?source_name=cute_ganymede&start_date=2024-11-14T22:05&end_date=2024-11-14T22:10"
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8082/logtracker/fetch_entries?source_name=cute_ganymede&start_date=2024-11-14T22:05&end_date=2024-11-14T22:10")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Response body: %s\n", string(bodyBytes))
}
