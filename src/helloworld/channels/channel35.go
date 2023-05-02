// https://habr.com/ru/articles/490336/
package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() { // enter the main goroutine
	fmt.Println("main() started") // print text
	c := make(chan string) // creating the new channel of type string

	go greet(c) // pass c (cannel of type string) to greet func and run greet() in goroutine

	// now greet running somwhere in goroutine
	
	// ? main goroutine is blocked before receiver gets the value

	c <- "John" // pass the string to the channel
	

	fmt.Println("main() stopped")
}