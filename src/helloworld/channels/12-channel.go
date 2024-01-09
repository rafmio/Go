// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
)

func squares(ch chan int) {
	for i := 0; i <= 9; i++ {
		ch <- i * i
	}

	close(ch)
}

func main() {
	fmt.Println("main() started")
	ch := make(chan int)

	go squares(ch)

	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("main() stopped")
}
