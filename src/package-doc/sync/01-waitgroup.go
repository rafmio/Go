package main

import (
	"fmt"
	"sync"
)

type CounterI struct {
	value int
}

func (c *CounterI) Update(n int, wg *sync.WaitGroup) {
	// fmt.Println("Enter the c.Update()")
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
}

func main() {
	var wg sync.WaitGroup
	c := CounterI{}

	valArr := [10]int{10, -5, 25, 19, 11, -10, 20, -9, 30, -1}
	wg.Add(len(valArr))

	fmt.Println("c.Value:", c.value)

	for _, val := range valArr {
		// fmt.Println(i, ":", val)
		go c.Update(val, &wg)
	}

	wg.Wait()
	fmt.Println("Final c.Value:", c.value)
}
