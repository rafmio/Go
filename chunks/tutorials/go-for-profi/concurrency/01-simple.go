package main

import (
	"fmt"
	"time"
)

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func main() {
	go function()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	time.Sleep(time.Millisecond * 300)
	fmt.Println()
}
