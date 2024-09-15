// https://habr.com/ru/articles/840750/
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go greeting(i)
	}
	time.Sleep(time.Second * 1)
}

func greeting(n int) {
	fmt.Printf("Hello from goroutine %d\n", n)
}
