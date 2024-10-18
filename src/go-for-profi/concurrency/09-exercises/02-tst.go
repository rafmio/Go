package main

import (
	"fmt"
	"sync"
)

func fillSls(wg *sync.WaitGroup, ch chan<- int) {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- i
			wg.Done()
		}()
	}
	fmt.Println("End of fillSls()")
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	sls := make([]int, 0)

	fillSls(&wg, ch)

	go func() {
		fmt.Println("anon func()")
		for num := range ch {
			sls = append(sls, num)
		}
	}()

	wg.Wait()

	fmt.Println(sls)
}
