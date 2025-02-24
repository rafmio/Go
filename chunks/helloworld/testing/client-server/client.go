package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func runClient() string {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// io.Copy(os.Stdout, resp.Body)

	defer resp.Body.Close()

	// fmt.Println(resp.Body)

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bodyStr := String(bodyByte)

	fmt.Println(bodyStr)

	return bodyStr
}

func String(bodyByte []byte) string {
	return fmt.Sprintf("%v", string(bodyByte))
}
