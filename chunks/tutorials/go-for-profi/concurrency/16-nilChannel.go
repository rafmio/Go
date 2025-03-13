package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c <-chan int) {
	fmt.Println("Starting add()...")
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C: // blocks channel 'C' on 't' time
			c = nil
			fmt.Println("Sum:", sum)
		}
	}
}

func send(c chan<- int) {
	fmt.Println("Starting send()...")
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	fmt.Println("Starting main()...")
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(time.Second * 3)
}
