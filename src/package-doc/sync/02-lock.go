package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	c.m.Lock()
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func main() {
	var wg sync.WaitGroup
	c := Counter{}

	sls := []int{10, -5, 25, 19, 1000}
	wg.Add(len(sls))

	for _, s := range sls {
		go c.Update(s, &wg)
	}

	wg.Wait()
}
