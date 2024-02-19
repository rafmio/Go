package main

import (
	"bytes"
	"fmt"
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
		log.Fatal("Error crearing new request:", err.Error())
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Huggy-Wuggy", "Kissy-Missy")
	req.Header.Set("Chikey-Pickey", "Tosy-Bosy")

	client := http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error doing GET http-request", err.Error())
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Proto:", resp.Proto)
	for key, value := range resp.Header {
		fmt.Println(key, ":", value)
	}
	_, err = io.Copy(os.Stdout, body)
	if err != nil {
		log.Fatal("Error copying to os.Stdout:", err.Error())
	}
}
