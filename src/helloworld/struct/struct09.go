// Empty struct
// https://golangbyexample.com/empty-struct-in-go/
package main

import (
  "fmt"
  "time"
)

func main() {
  done := make(chan struct{}, 4)
  for i := 0; i < 4; i++ {
    go runasync(done)
  }

  for i := 0; i < 4; i++ {
    fmt.Println(<-done)
  }
  close(done)
  fmt.Printf("Finish\n")
}

func runasync(done chan<- struct{}) {
  time.Sleep(time.Second)
  done <-struct{}{}
  return
}
