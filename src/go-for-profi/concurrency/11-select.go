package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(time.Second * 4):
			fmt.Println("Generating numbers took too long. Stopping.")
		}
	}
}

func main() {
	createNumber := make(chan int)
	end := make(chan bool)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [max_number]")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid max number: %s\n", os.Args[1])
		return
	}

	fmt.Printf("Going to create %d random numbers\n", n)
	go gen(0, 2*n, createNumber, end)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Exiting...")
	end <- true
}
