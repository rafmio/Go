// https://pkg.go.dev/time#After
// select, time.After()
package main

import (
  "fmt"
  "time"
)

var c chan int
// c := make(chan int) - equivalent for previous inside main func 

func handle(int) {}

func main() {

  select {
  case m := <-c:
    handle(m)
  case <- time.After(4 * time.Second):
    fmt.Println("timed out")
  }
}
