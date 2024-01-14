// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	time.Sleep(time.Second * 3)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(time.Second * 4)
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	select {
	case res := <-chan1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response from servise 2", res, time.Since(start))
	case <-time.After(time.Second * 2):
		fmt.Println("No response received", time.Since(start))
	}

	close(chan1)
	close(chan2)

	fmt.Println("main() stopped", time.Since(start))
}
