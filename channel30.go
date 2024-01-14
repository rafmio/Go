// https://habr.com/ru/post/490336/
package main

import (
  "fmt"
  "sync"
  "time"
  "runtime"
)

func sqrWorker(wg *sync.WaitGroup, tasks <-chan int, result chan<- int,
instance int) {
  for num := range tasks {
    time.Sleep(time.Millisecond)
    fmt.Printf("[worker %v] Sending result by worker %v. Sends num: %v\n", instance, instance, num)
    result <- num * num
  }
  wg.Done() // done with worker
}

func main() {
  fmt.Println("[main] main() started")

  var wg sync.WaitGroup

  tasks := make(chan int, 10)
  results := make(chan int, 10)

  fmt.Println("[main] active goroutenes: ", runtime.NumGoroutine())

  // launching 3 worker goroutines
  for i := 0; i < 3; i++ {
    wg.Add(1)
    go sqrWorker(&wg, tasks, results, i)
  }
  fmt.Println("[main] active goroutenes: ", runtime.NumGoroutine())
  fmt.Println()

  // passing 5 tasks
  for i := 0; i < 5; i++ {
    tasks <- i * 2  //  non-blocking as buffer capacity is 10
  }

  fmt.Println("[main] Wrote 5 tasks")

  // closing tasks
  close(tasks)

  // wait until all workers done their job
  wg.Wait()


  fmt.Println("[main] active goroutenes: ", runtime.NumGoroutine())
  fmt.Println()
  
  // receiving results from all workers
    for i := 0; i < 5; i++ {
      result := <- results  // non-blocking because buffer is non-empty
      fmt.Println("[main] Result", i, ":", result)
    }


  fmt.Println("[main] active goroutenes: ", runtime.NumGoroutine())
  fmt.Println()

  fmt.Println("[main] main() stopped")
}
