// https://habr.com/ru/articles/490336/
package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 1000; i++ {
		c <- i * i
	}

	close(c)
}

func main() {
	runChannel()
}

func runChannel() {
	fmt.Println("runChannel() started")
	c := make(chan int, 16)

	go squares(c)

	for val := range c {
		fmt.Printf("%d ", val)
	}

	fmt.Println("main() stopped")
}