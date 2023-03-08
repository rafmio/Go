package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Println("Working...")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()

	wg.Wait()
}
