package main

import (
	"fmt"
	"time"
)

func producer(items chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Producing item %d\n", i)
		items <- i
	}

	close(items)
}

func consumer(items <-chan int) {
	for item := range items {
		fmt.Printf("Consuming item %d\n", item)
	}
}

func main() {
	items := make(chan int)
	go producer(items)
	go consumer(items)

	time.Sleep(1000 * time.Millisecond)
}
