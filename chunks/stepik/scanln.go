package main

import "fmt"

func main() {
  var input int
  fmt.Scanln(&input)
  fmt.Println("input * 2", input * 2)

  var a, b, c int
  fmt.Scanln(&a, &b, &c)
  fmt.Println("a = ", a, "b = ", b, "c = ", c)
}
