// https://habr.com/ru/post/490336/

package main

import (
	"fmt"
)

func great(c chan string) {
	<-c
	<-c
}

func main() {
	fmt.Println("main() started")
	c := make(chan string, 1)

	go great(c)

	c <- "John"

	// close(c)

	c <- "Mike"

	fmt.Println("main() stopped")
}
