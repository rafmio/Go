// Converting a string to runes
// rune = Unicode point (may be more than one byte) = singe character
package main

import (
  "fmt"
  "strconv"
)
func main() {
  var price []rune = []rune("€48.95")
  var currency string = string(price[0])
  var amountString string = string(price[1:])
  amount, parseErr := strconv.ParseFloat(amountString, 64)

  fmt.Println("Length:", len(price))
  fmt.Println("Currency:", currency)
  if (parseErr == nil) {
    fmt.Println("Amount:", amount)
  } else {
    fmt.Println("Parse Error:", parseErr)
  }

  var price2 = "€17.25 ЖИ"
  for index, char := range price2 {
    fmt.Println(index, char, string(char), rune(char))
  }
  fmt.Println()
  for index, char := range []byte(price2) {
    fmt.Println(index, char)
  }
}
