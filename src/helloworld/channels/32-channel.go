// https://habr.com/ru/post/490336/
// Mutex

package main

import (
	"fmt"
	"sync"
)

var i int // i == 0

// goroutine increment global variable i
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	// wait until all 1000 goroutines are done
	wg.Wait()

	fmt.Println("value of i after 1000 operations is:", i)
}
