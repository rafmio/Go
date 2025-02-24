package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64
	counter = 0

	// Атомарное увеличение счетчика
	returnedVal := atomic.AddInt64(&counter, 1)
	fmt.Printf("Counter after increment: %d, returnedVad: %d\n", counter, returnedVal)

	returnedVal = atomic.AddInt64(&counter, 1)
	fmt.Printf("Counter after increment: %d, returnedVad: %d\n", counter, returnedVal)

	time.Sleep(time.Millisecond * 300)
}
