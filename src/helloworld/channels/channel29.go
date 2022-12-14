// https://habr.com/ru/post/490336/
// "Пул воркеров"
package main

import (
	"fmt"
	"time"
)

func sqrWorker(tasks <-chan int, results chan<- int, id int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking tasks
		fmt.Printf("[worker %v] Sending result by worker %v, num: %v\n", id, id, num)
		results <- num * num
	}
}

func main() {
	fmt.Println("[main] main() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receiving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("[main] main() stopped")
}
