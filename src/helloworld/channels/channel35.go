// https://habr.com/ru/articles/490336/
package main

import (
	"fmt"
	"sync"
)

var j int 

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	j = j + 1
	m.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()

	fmt.Println("value of j after 1000 operations is: ", j)
}