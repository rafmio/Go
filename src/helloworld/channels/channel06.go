package main

import "fmt"

func sending(s chan<- string) {
	s <- "Seek and destroy"
}

func main() {
	mychanl := make(chan string)
	go sending(mychanl)
	fmt.Println(<-mychanl)
}

// Unidirectional channel
// https://www.geeksforgeeks.org/unidirectional-channel-in-golang/
