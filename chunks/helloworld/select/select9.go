// https://www.geeksforgeeks.org/time-after-function-in-golang-with-examples/
package main

import (
  "fmt"
  "time"
)

var ch chan int

func main() {
  for i := 1; i < 6; i++ {
    fmt.Println(i, " Welcome, Greatest Go Developer!")
  }

  select {
  case <- time.After(3 * time.Second):
    fmt.Println("Time out")
  }
}
