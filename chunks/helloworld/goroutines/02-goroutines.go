package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(w, str)
	}
}

func main() {
	go display("Welcome!")
	display("Come here to me")
}
