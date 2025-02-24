// https://pkg.go.dev/sync#WaitGroup
package main

import (
	"fmt"
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
		"http://www.yandex.ru/",
		"http://www.stackoverflow.com/",
	}

	for _, url := range urls {
		// Increment the WaitGroup counter
		wg.Add(1)

		// Launch a goroutine to fetch the URL
		go func(url string) {
			// Decrement the counter when the goroutine completes
			defer wg.Done()
			// Fetch the URL
			response, err := http.Get(url)
			if err != nil {
				fmt.Println("ERROR: http.Get()", err)
			}
			fmt.Printf("Fetched %s: status %d, proto: %s\n", url, response.StatusCode, &response.Proto)
		}(url)
	}

	// Wait for all the goroutines to complete
	wg.Wait()
}
