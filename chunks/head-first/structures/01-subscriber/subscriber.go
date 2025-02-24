package main

import (
  "fmt"
  "magazine"
)

func main() {
  var s magazine.subscriber
  s.rate = 4.99
  fmt.Println(s.rate)
}
