// https://habr.com/ru/post/490336/
package main

import "fmt"
import "time"

func main() {
  fmt.Println("main() started")
  c := make(chan string)

  go func(c chan string) {
    fmt.Println("Hello " + <-c + "!")
  } (c)

  c <- "John"
  fmt.Println("main() stopped")

  go func() {
    fmt.Println("Anonymous func")
  } ()


  anfun := func() {
    fmt.Println("Init anfun")
  }
  go anfun()

  time.Sleep(time.Second * 1)
}
