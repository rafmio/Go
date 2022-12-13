// https://golangify.com/goroutines
package main

import (
  "fmt"
  "time"
)

func main() {
  c := make(chan int)
  
  for i := 0; i < 5; i++ {
    go sleepyGopher(i, c)
  }
  for i := 0; i < 5; i++ {
    gopherID := <- c
    fmt.Println("gopher ", gopherID, " has finished sleeping")
  }
}

func sleepyGopher(id int, c chan int) {
  fmt.Println("Process ", id, "doing smth 3 seconds...")
  time.Sleep(time.Second * 3)
  fmt.Println("...", id, "snore...")
  c <- id
}
