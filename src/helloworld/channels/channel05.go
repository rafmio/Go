package main

import "fmt"

func main() {
	mychan1 := make(<-chan string)
	mychan2 := make(chan<- string)

	fmt.Printf("%T", mychan1)
	fmt.Printf("\n%T", mychan2)
	fmt.Println()
}

// Unidirectional channel
// https://www.geeksforgeeks.org/unidirectional-channel-in-golang/
