// https://habr.com/ru/post/490336/
package main

import "fmt"

func greet(c chan string) {
  fmt.Println("Hello " + <- c + "!")
  defer close(c)
}

func main() {
  fmt.Println("main() started")
  c := make(chan string)

  go greet(c)

  c <- "John"
  fmt.Println("main() stopped")
}
