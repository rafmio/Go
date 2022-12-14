// https://habr.com/ru/post/490336/
package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2

	fmt.Printf("len(c): %v, cap(c): %v", len(c), cap(c))
	fmt.Println()

	fmt.Println("c:", <-c)
	fmt.Println("c:", <-c)
}
