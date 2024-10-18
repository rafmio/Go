package main

import "sync"

func fillSls(wg *sync.WaitGroup, ch chan<- int) {
	wg.Add(1)
	defer wg.Done()

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

}
