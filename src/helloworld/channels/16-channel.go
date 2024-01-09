// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
	"os"
	"runtime"
)

func squares(ch chan int) {
	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started, PID:", os.Getpid())
	ch := make(chan int, 3)

	go squares(ch)

	fmt.Println("active goroutines", runtime.NumGoroutine())
	for i := 1; i <= 4; i++ {
		ch <- i
	}

	fmt.Println("active goroutines", runtime.NumGoroutine())

	go squares(ch)

	fmt.Println("active goroutines", runtime.NumGoroutine())

	for i := 5; i <= 9; i++ {
		ch <- i
	}

	fmt.Println("active goroutines", runtime.NumGoroutine())

	fmt.Println("main() stopped")
}
