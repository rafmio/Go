package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer!")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println(err)
		return
	}

	fmt.Printf("Going to create %d random numbers\n", n)

	go monitor()

	var wg sync.WaitGroup

	for r := 0; r < n; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(n*10 + r))
		}()
	}
	wg.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}
