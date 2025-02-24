package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	userAgent := "Huizilla/1.0, Huindows x86_64 AppleWebKit"
	var buffer bytes.Buffer
	body := io.Reader(&buffer)

	req, err := http.NewRequest("GET", "http://194.58.102.129:9001/", body)
	if err != nil {
		log.Fatal("ERROR creating new request")
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Arbitrary-header", "header-huyeader")

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("ERROR reading response:", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, body)
	if err != nil {
		log.Fatal("ERROR copying in os.Stdout:", err)
	}
}
