package main

import "fmt"

func boring(msg string) <-chan string { //Returns receive-only channel of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c // Return the channel to the caller.
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring; I'm leaving")
}
