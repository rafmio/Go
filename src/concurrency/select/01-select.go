package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go One(ch1)
	go Two(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println("Message from channel 1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Message from channel 2:", msg2)
	}
}

func One(ch chan<- string) {
	ch <- "Hello from channel 1"
}

func Two(ch chan<- string) {
	ch <- "Hello from channel 2"
}
