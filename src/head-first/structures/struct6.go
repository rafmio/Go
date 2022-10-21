package main

import "fmt"

type subscriber struct {
  name string
  rate float64
  active bool
}

func main() {
  var subscriber1 subscriber
  subscriber1.name = "Aman Singh"
  subscriber1.active = true
  fmt.Println("Name: ", subscriber1.name)

  var subscriber2 subscriber
  subscriber2.name = "Beth Ryan"
  subscriber2.rate = 34.43
  subscriber2.active = true
  fmt.Println("Name: ", subscriber2.name, "| Rate: ", subscriber2.rate,
              "| Active? :", subscriber2.active)
}
