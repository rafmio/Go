// https://habr.com/ru/post/490336/

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2

	fmt.Printf("len(ch): %v, cap(ch): %v\n", len(ch), cap(ch))

	fmt.Println("ch:", <-ch)
	fmt.Println("ch:", <-ch)
}
