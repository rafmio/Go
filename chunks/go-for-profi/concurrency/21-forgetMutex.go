package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

func function() {
	m.Lock()
	fmt.Println("Locked!")
	// m.Unlock()

}

func main() {
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		function()
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()
		function()
	}()

	wg.Add(1)

	wg.Wait()
}
