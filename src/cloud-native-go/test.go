package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var num int

func responseSize(url string) {
	defer wg.Done()

	num++
	fmt.Println(num, "Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	num++
	fmt.Println(num, "Step2: ", url)
	defer response.Body.Close()

	num++
	fmt.Println(num, "Step3: ", url)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	num++
	fmt.Println(num, "Step4: len(body)", len(body))
}

func main() {
	// wg.Add(3)
	fmt.Println("Start Goroutines")

	go responseSize("https://rafmio.ru")
	wg.Add(1)
	go responseSize("http://himpostavka.ru")
	wg.Add(1)
	go responseSize("https://ya.ru")
	wg.Add(1)

	wg.Wait()
	fmt.Println("Terminated Program")
}
