// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
)

// fib returns a channel which transports fibonacci numbers
func fib(length int) <-chan int {
	// make buggered channel
	c := make(chan int)

	// run generation concurently
	go func() {
		for i, j := 0, 1; i < length; i, j = i+j, i {
			c <- i
		}
		close(c)
	}()

	return c
}

func main() {
	// read 10 fibonacci numbers from channel returned by `fib` func
	for fn := range fib(10) {
		fmt.Println("Current fibonacci number is:", fn)
	}
}
