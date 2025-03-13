package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func change(i int) {
	m.Lock()
	time.Sleep(time.Millisecond * 100)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
		// fmt.Printf("i: %d, v1%10 == 0, v1%10 == %d\n", i, v1%10)
	} else {
		// fmt.Printf("i: %d, v1%10 != 0, v1%10 != %d\n", i, v1%10)
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	// fmt.Printf("\nType of a: %T, type of v1: %T\n", a, v1)
	m.Unlock()
	return a
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup

	fmt.Printf("%d ", read())
	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf(">- %d\n", read())
		}(i)
	}

	waitGroup.Wait()
	fmt.Printf("-> %d\n", read())
}
