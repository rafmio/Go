package main

import (
	"fmt"
	"sync"
)

func sendToChan(ch chan int, val int) {
	ch <- val
}

func main() {
	ch := make(chan int)
	defer close(ch)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		go sendToChan(ch, i*100)
		wg.Add(1)
	}

	for i := 1; i <= 5; i++ {
		go func() {
			fmt.Printf("%d ", <-ch)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println()
}
