// https://habr.com/ru/post/490336/
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, result chan<- int, instance int) {
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] sending result by worker %v. Sends num: %v\n",
			instance, instance, num)
		result <- num * num
	}

	wg.Done()
}

func main() {
	fmt.Println("[main] main() started")

	var wg sync.WaitGroup

	tasks := make(chan int, 10)
	results := make(chan int, 10)
	// defer close(tasks)
	// defer close(results)

	fmt.Println("[main] active goroutines:", runtime.NumGoroutine())

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqrWorker(&wg, tasks, results, i)
		fmt.Printf("worker %d created and added to WaitGroup\n", i)
		time.Sleep(time.Second * 1)
	}

	fmt.Println()
	fmt.Println("[main] active goroutines:", runtime.NumGoroutine())
	fmt.Println()

	// passing 5 tasks:
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
		fmt.Printf("[main] value [%d] passed to the 'tasks' channel --- ", i)
		time.Sleep(time.Second * 1)
	}

	fmt.Println("[main] write 5 tasks")

	close(tasks)

	// wait until all worker done their job
	wg.Wait()

	fmt.Println("[main] active goroutines:", runtime.NumGoroutine())
	fmt.Println()

	// receiving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // non-blocking because buffer is non-empty
		fmt.Println("[main] result", i, ":", result)
	}

	fmt.Println("[main] active goroutines:", runtime.NumGoroutine())
	fmt.Println()

	fmt.Println("[main] main() stopped")
}
