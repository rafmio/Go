package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Timeout finished")

	ch <- 100
}

func receive(ch <-chan int) {
	val := <-ch
	fmt.Println("Receiving value from channel: ", val)
}

func main() {
	ch := make(chan int)

	go send(ch)
	go receive(ch)

	time.Sleep(time.Second * 2)
}

// https://golangbyexample.com/channel-golang/
