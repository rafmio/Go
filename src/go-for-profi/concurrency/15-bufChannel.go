package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int, 7)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println("Received:", num)
		default:
			fmt.Println("No more numbers available")
			break
		}
	}
}
