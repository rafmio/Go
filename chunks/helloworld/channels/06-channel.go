package main

import (
	"fmt"
)

func sending(s chan<- string) {
	s <- "Seek and destroy"
}

func main() {
	mychan := make(chan string)
	go sending(mychan)
	fmt.Println(<-mychan)
}

// Unidirectional channel
// https://www.geeksforgeeks.org/unidirectional-channel-in-golang/
