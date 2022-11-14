package main

import (
  "fmt"
)

func main() {
  product := "Kayak"
  for index, character := range product {
    fmt.Println("Index:", index, "Charhacter:", string(character))
  }
  fmt.Printf("Type of product: %T\n", product)
  fmt.Println()

  for index := range product {
    fmt.Println("Index:", index)
  }
  fmt.Println()

  for _, character := range product {
    fmt.Println("Charhacter:", string(character))
  }
  fmt.Println()

  products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
  for index, element := range products {
    fmt.Println("Index:", index, "Element:", element)
  }
  fmt.Println()
}
