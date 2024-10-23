package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// atomic.AddInt64(&counter, int64(i)) // will be only one result every run
			counter += int64(i) // will lead to different results every run
		}()
	}

	wg.Wait()
	value := atomic.LoadInt64(&counter)
	fmt.Printf("Counter value: %d\n", value)
}
