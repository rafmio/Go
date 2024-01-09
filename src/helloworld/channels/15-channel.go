// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	defer close(ch)
	for i := 1; i <= 6; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int, 3)

	go sender(ch)

	fmt.Printf("len(ch): %v, cap(ch): %v\n", len(ch), cap(ch))
	time.Sleep(time.Second * 2)
	fmt.Printf("len(ch): %v, cap(ch): %v\n", len(ch), cap(ch))

	for val := range ch {
		fmt.Printf("val: %v, len(ch): %v, cap(ch): %v\n", val, len(ch), cap(ch))
	}

	fmt.Println()
}
