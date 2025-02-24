// Sorting slices
package main

import (
  "fmt"
  "sort"
)

func main() {
  products := []string { "Kayak", "Lifejacket", "Paddle", "Hat" }
  fmt.Println(products)
  sort.Strings(products)
  fmt.Println(products)
  fmt.Println()

  for index, value := range products {
    fmt.Println("Index:", index, "value:", value)
  }
}
