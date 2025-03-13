package main

import "fmt"

type Number int

func (n *Number) Double() {
  *n *= 2
}

func main() {
  number := Number(4)
  fmt.Println("Original value of number: ", number)
  number.Double()
  fmt.Println("number after calling Double: ", number)

  var newNumber Number = 10
  newNumber.Double()
  fmt.Println("newNumber: ", newNumber)

}
