// https://www.geeksforgeeks.org/time-after-function-in-golang-with-examples/
package main

import (
  "fmt"
  "time"
)

func main() {
  channel := make(chan string, 2)

  select {
  case output := <- channel:
    fmt.Println(output)
  case <- time.After(4 * time.Second):
    fmt.Println("It's timeout...")
  }
}
