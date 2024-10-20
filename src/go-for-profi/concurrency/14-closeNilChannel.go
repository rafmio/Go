package main

import (
	"fmt"
)

func main() {
	fmt.Println("Closing nil channel")
	var c chan string
	close(c) // Will lead to a 'panic'
}
