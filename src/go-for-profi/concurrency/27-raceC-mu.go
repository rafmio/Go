package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func writeMap(k map[int]int, key, value int) {
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()
	k[key] = value
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	var wg sync.WaitGroup

	var i int

	k := make(map[int]int)
	var mu sync.Mutex

	// k[1] = 12
	mu.Lock()
	writeMap(k, 1, 12)
	mu.Unlock()

	for i = 0; i < numGR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// k[i] = i
			mu.Lock()
			writeMap(k, i, i)
			mu.Unlock()
		}(i)
	}

	// k[2] = 10
	mu.Lock()
	writeMap(k, 2, 10)
	mu.Unlock()

	wg.Wait()
	fmt.Printf("k = %#v\n", k)
}
