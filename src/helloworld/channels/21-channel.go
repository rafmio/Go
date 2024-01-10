// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main() started")
	ch := make(chan string)

	go func(ch chan string) {
		fmt.Println("Hello " + <-ch + "!")
	}(ch)

	ch <- "Will"
	fmt.Println("main() stopped")

	go func() {
		fmt.Println("Anonymous func")
	}()

	anfun := func() {
		fmt.Println("Init anfun")
	}

	go anfun()

	time.Sleep(time.Second * 1)
}
