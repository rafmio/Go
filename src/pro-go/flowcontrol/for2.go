package main

import (
  "fmt"
)

func main() {
  for counter := 0; counter <= 3; counter++ {
    if (counter == 1) {
      continue
      fmt.Println("Current counter's value:", counter) // Never been printed
    }
    fmt.Println("Counter:", counter)
  }
}
