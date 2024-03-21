package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	// rand.NewSource(time.Now().UnixNano())

	numbers := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numbers <- rand.Intn(100)
		}
		close(numbers)
	}()

	for n := range numbers {
		fmt.Printf("%d ", n)
	}

	fmt.Println()
}
