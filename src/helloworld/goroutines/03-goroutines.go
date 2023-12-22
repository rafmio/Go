package main

import (
	"fmt"
	"os"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println("PID: ", os.Getpid(), "String inside for loop: ", str)
	}
}

func main() {
	go display("Hello inside Goroutine")
	display("Hole inside normal function")
}
