package main

import (
	"fmt"
)

func boring(msg string) <-chan string { //Returns receive-only channel of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c // Return the channel to the caller.
}

func main() {
	bc := boring("Greetings")
	for i := 0; i < 5; i++ {
		fmt.Println(<-bc) // Receive from channel and print the received string
	}
	bc = boring("Goodbye")
	for i := 0; i < 5; i++ {
		fmt.Println(<-bc) // Receive from channel and print the received string
	}

	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println("joe:", <-joe)
		fmt.Println("ann:", <-ann)
	}

	fmt.Println("You're both boring; I'm leaving")
}
