// https://habr.com/ru/post/490336/
package main

import (
  "fmt"
  "time"
)

var start time.Time

func init() {
  start = time.Now()
}

func main() {
  fmt.Println("main() started", time.Since(start))
  chan1 := make(chan string, 2)
  chan2 := make(chan string, 2)

  chan1 <- "Value 1"
  chan1 <- "Value 2"
  chan2 <- "Value 3"
  chan2 <- "Value 4"

  select {
  case res := <- chan1:
    fmt.Println("Response from chan1", res, time.Since(start))
  case res := <- chan2:
    fmt.Println("Response from chan2", res, time.Since(start))
  }

  fmt.Println("main() stopped", time.Since(start))
}
