// https://habr.com/ru/post/490336/
package main

import (
  "fmt"
  "runtime"
)

func squares(c chan int) {
  for i := 0; i < 4; i++ {
    num := <- c
    fmt.Println(num * num)
  }
}

func main() {
  fmt.Println("main() started")
  c := make(chan int, 3)

  go squares(c)

  fmt.Println("active goroutines", runtime.NumGoroutine())
  c <- 1
  c <- 2
  c <- 3
  c <- 4 // block here

  fmt.Println("active goroutines", runtime.NumGoroutine())

  go squares(c)

  fmt.Println("active goroutines", runtime.NumGoroutine())

  c <- 5
  c <- 6
  c <- 6
  c <- 8
  c <- 9

  fmt.Println("active goroutines", runtime.NumGoroutine())
  fmt.Println("main() stopped")
}
