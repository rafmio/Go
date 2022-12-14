// https://habr.com/ru/post/490336/
package main

import (
  "fmt"
  "sync"
  "time"
)

func service(wg *sync.WaitGroup, instance int) {
  time.Sleep(2 * time.Second)
  fmt.Println("Service called on instance", instance)
  wg.Done() // Decrement counter
}

func main() {
  fmt.Println("main() started")

  var wg sync.WaitGroup // Create waitgrout (empty struct)

  for i := 0; i <= 3; i++ {
    wg.Add(1) // Increment counter
    go service(&wg, i)
  }

  wg.Wait() // Block here
  fmt.Println("main() stopped")
}
