package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:9000"

	var path string
	fmt.Printf("Enter the path: ")
	fmt.Scanf("%s", &path)
	if path != "" {
		url = url + "/" + path
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("ERROR creating request:", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERROR sending request:", err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("ERROR reading response body", err)
	}

	fmt.Printf("Response: %s\n", string(body))
}
