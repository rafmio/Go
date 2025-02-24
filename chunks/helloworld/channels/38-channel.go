// https://golangbyexample.com/channel-golang/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendFunc(ch chan int) {
	for i := 0; i < 15; i++ {
		time.Sleep(time.Millisecond * 10)
		ch <- rand.Intn(50)
	}
}

func receiveFunc(ch <-chan int) {
	for val := range ch {
		fmt.Printf("%d ", val)
	}
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go sendFunc(ch)
	go receiveFunc(ch)

	time.Sleep(time.Second * 1)
	fmt.Println()
}
