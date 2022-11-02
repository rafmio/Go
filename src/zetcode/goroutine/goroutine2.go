package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  var wg sync.WaitGroup
  fmt.Printf("wg value: %v, wg type: %T\n", wg, wg)
  wg.Add(2)
  fmt.Println("After wg.Add(2): ")
  fmt.Printf("wg value: %d, wg type: %T\n", wg, wg)

  go func() {
    count("oranges")
    wg.Done()
  }()

  go func() {
    count("apples")
    wg.Done()
  }()

  wg.Wait()
  fmt.Printf("wg value: %d, wg type: %T\n", wg, wg)
}

func count(thing string) {
  for i := 0; i < 4; i++ {
    time.Sleep(time.Millisecond * 500)
    fmt.Printf("counting %s\n", thing)
  }
}
