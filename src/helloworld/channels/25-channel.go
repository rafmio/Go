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
	fmt.Println("Service 1 started:", time.Since(start))
	c <- "Hello from service 1"
}

func service2(c chan string) {
	fmt.Println("Service 2 started:", time.Since(start))
	c <- "Hello from service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(chan2)

	// if comment below - the result will change
	// time.Sleep(time.Second * 1)

	select {
	case res := <-chan1:
		fmt.Println("Response form service 1", res, time.Since(start))
	case res := <-chan2:
		fmt.Println("Response form service 2", res, time.Since(start))
	default:
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
}
