// Dual nature of stirngs
package main

import (
  "fmt"
  "strconv"
)

func main() {
  var price string = "$48.95"
  for i, char := range price {
    fmt.Println(i, string(char))
  }
  fmt.Println()

  var currency byte = price[0]
  fmt.Println("currency:", currency)
  fmt.Println("currency:", string(currency))
  fmt.Println()

  var amountString string = price[1:]
  amount, parseErr := strconv.ParseFloat(amountString, 64)

  fmt.Println("Currency:", currency)
  if (parseErr == nil) {
    fmt.Println("Amount:", amount)
    fmt.Printf("Type of amount: %T\n", amount)
  } else {
    fmt.Println("Parse Error:", parseErr)
  }
}
