package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

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
	k[1] = 12

	for i = 0; i < numGR; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			k[i] = i
		}()
	}

	k[2] = 10
	wg.Wait()
	fmt.Printf("k = %#v\n", k)
}

// go run -race 26-raceC.go 10
