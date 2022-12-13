package main

import (
  "fmt"
  "time"
)

func main() {
  for i := 0; i < 5; i++ {
    go sleepyGopher()
    // time.Sleep(time.Second * 2)
  }
  time.Sleep(time.Second * 4)
}

func sleepyGopher() {
  time.Sleep(time.Second * 1)
  fmt.Println(" ... snore ... ")
}
