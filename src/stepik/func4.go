package main

import "fmt"

func double_m(a, b int) int {
  var total int
  for i := a; i <= b; i++ {
    total += i * i
    fmt.Println(total)
  }
  return total
}

func main() {
  fmt.Println(double_m(2, 6))
}
