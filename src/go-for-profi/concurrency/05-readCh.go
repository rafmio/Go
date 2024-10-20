package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)

	go writeToChannel(c, 10)

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Read:", <-c)
	time.Sleep(time.Millisecond * 100)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
