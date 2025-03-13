package main

import "fmt"

type subscriber struct {
  name string
  rate float64
  active bool
}

func applyDiscounts(s *subscriber) {
  s.rate = 4.99
  s.active = true
}

func main() {
  var s subscriber
  applyDiscounts(&s)
  fmt.Println("s.rate: ", s.rate)
  fmt.Println("s.active: ", s.active)
}
