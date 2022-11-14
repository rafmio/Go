package main

import (
  "fmt"
)

func main() {
  kayakPrice := 275.00
  if kayakPrice > 100 {
    fmt.Println("Price is greater than 100")
  }
  if (kayakPrice < 400) {
    fmt.Println("Price is less than 400")
    fmt.Println("Go doesn't require parentheses for condition")
  }
  fmt.Println()

  if (kayakPrice > 500) {
    fmt.Println("Price is greater than 500")
  } else if (kayakPrice < 300) {
    fmt.Println("Price is less than 300")
  }
}
