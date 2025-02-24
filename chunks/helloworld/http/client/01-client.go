// https://networkbit.ch/golang-http-client/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.example.com")
	if err != nil {
		log.Fatal("ERROR getting response: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ERROR reading response: ", err)
	}

	err = os.WriteFile("body.html", body, 0644)
	if err != nil {
		fmt.Println("ERROR saving to file: ", err)
	}

	for key, value := range resp.Header {
		for k, v := range value {
			fmt.Printf("[%s] - [%v] : %s\n", key, k, v)
		}
	}
}
