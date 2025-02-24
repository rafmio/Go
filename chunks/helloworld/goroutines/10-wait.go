// https://habr.com/ru/articles/840750/
package main

import (
	"fmt"
	"sync"
)

func main() {
	// creating the WaitGroup
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// time.Sleep(time.Duration(i) * time.Millisecond)
			fmt.Println("Goroutine", i, "has been finished it's execution")
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines has been finished")
}
