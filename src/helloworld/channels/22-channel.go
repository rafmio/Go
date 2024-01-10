// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
)

func greet(ch chan string) {
	fmt.Println("Hello " + <-ch + "!")
}

func greeter(cc chan chan string) {
	ch := make(chan string)
	cc <- ch
}

func main() {
	fmt.Println("main() started")

	cc := make(chan chan string)

	go greeter(cc)

	c := <-cc

	go greet(c)

	c <- "Will"

	fmt.Println("main() stopped")
}
