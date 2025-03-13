package main

import "fmt"

type subscriber struct {
  name string
  rate float64
  active bool
}

func printInfo(s *subscriber) {
  fmt.Println("Name: ", s.name)
  fmt.Println("Rate: ", s.rate)
  fmt.Println("Active?: ", s.active)
}

func defaultSubscriber(name string) *subscriber {
  var s subscriber
  s.name = name
  s.rate = 5.99
  s.active = true
  return &s
}

func applyDiscount(s *subscriber) {
  s.rate = 4.99
}

func main() {
  subscriber1 := defaultSubscriber("John Lennon")
  applyDiscount(subscriber1)
  printInfo(subscriber1)

  subscriber2 := defaultSubscriber("Paul McCartney")
  printInfo(subscriber2)
}
