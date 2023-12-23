package main

import (
	"fmt"
	"os"
)

func main() {
	mychan1 := make(<-chan string)
	mychan2 := make(chan<- string)

	fmt.Printf("Type of mychan1: %T\n", mychan1)
	fmt.Printf("Type of mychan2: %T\n", mychan2)

	os.Exit(0)
}

// Unidirectional channel
// https://www.geeksforgeeks.org/unidirectional-channel-in-golang/
