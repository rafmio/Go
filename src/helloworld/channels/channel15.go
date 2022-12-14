// https://habr.com/ru/post/490336/
package main

import "fmt"
import "time"

func sender(c chan int) {
	defer close(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
}

func main() {
	c := make(chan int, 3)

	go sender(c)

	fmt.Printf("len(c): %v, cap(c): %v", len(c), cap(c))
	fmt.Println()
	time.Sleep(time.Second * 2)
	fmt.Printf("len(c): %v, cap(c): %v", len(c), cap(c))
	fmt.Println()

	for val := range c {
		fmt.Printf("val: %v, len(c): %v, cap(c): %v\n", val, len(c), cap(c))
	}
	fmt.Println()
}
