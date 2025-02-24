// https://golangify.com/goroutines
// Размер буфера канала

package main

import (
	"fmt"
	"time"
)

func squares(ch chan int) {
	for i := 0; i <= 3; i++ {
		num := <-ch
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	ch := make(chan int, 3)

	go squares(ch)

	ch <- 1
	ch <- 2
	ch <- 4
	ch <- 8

	defer time.Sleep(time.Second * 2)

	fmt.Println("main() stopped")
}
