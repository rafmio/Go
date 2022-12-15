// https://habr.com/ru/post/490336/
// Mutex
// The code below is with race condition
package main

import (
  "fmt"
  "sync"
  // "runtime"
)

var i int   // i == 0

// goroutine increment global variable i
func worker(wg *sync.WaitGroup) {
  i = i + 1
  wg.Done()
}

func main() {
  var wg sync.WaitGroup

  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go worker(&wg)
    // fmt.Println("active goroutines:", runtime.NumGoroutine())
  }

  wg.Wait()

  fmt.Println("value of i after 1000 operations:", i)
}
