package main

import "fmt"

func main() {
  var budget int
  fmt.Scanln(&budget)
  switch {
  case budget > 1000:
    fmt.Println("Apple")
  case budget >= 500 && budget <= 1000:
    fmt.Println("Samsung")
  case budget < 500:
    fmt.Println("Nokia с фонариком")

  }
}
