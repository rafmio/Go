package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func runClient() io.ReadCloser {
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	return resp.Body
}
