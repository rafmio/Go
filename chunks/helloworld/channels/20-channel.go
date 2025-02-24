// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
)

func greet(roc <-chan string) {
	fmt.Println("Hello " + <-roc + "!")
}

func main() {
	fmt.Println("main() started")

	var str string
	fmt.Printf("Enter the string: ")
	fmt.Scanf("%s", &str)

	ch := make(chan string)

	go greet(ch)

	ch <- str

	fmt.Println("main() stopped")
}
