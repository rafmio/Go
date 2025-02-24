// https://habr.com/ru/post/490336/

package main

import (
	"fmt"
)

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
	defer close(c)
}

func main() {
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- name

	fmt.Println("main() stopped")
}
