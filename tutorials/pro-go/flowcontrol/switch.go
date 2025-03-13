package main

import "fmt"

func main() {
  product := "Kayak"
  for index, character := range product {
    switch (character) {
    case 'K':
      fmt.Println("K at position", index)
    case 'y':
      fmt.Println("y at position", index)
    case 'a':
      fmt.Println("a at position", index)
    case 'k':
      fmt.Println("k at position", index)
    }
  }
  fmt.Println()

  for index, character := range product {
    switch (character) {
    case 'K', 'k':
      fmt.Println("K or k at position", index)
    case 'y':
      fmt.Println("y at position", index)
    }
  }
  fmt.Println()

  for index, character := range product {
    switch (character) {
    case 'K', 'k':
      if (character == 'k') {
        fmt.Println("Lowercase k at position", index)
        break
      }
      fmt.Println("Uppercase K at position", index)
    case 'y':
      fmt.Println("y at position", index)
    }
  }
  fmt.Println()

  for index, character := range product {
    switch (character) {
    case 'K':
      fmt.Println("Uppercase character")
      fallthrough
    case 'k':
      fmt.Println("k at position", index)
    case 'y':
      fmt.Println("y at position", index)
    }
  }
  fmt.Println()

  for index, character := range product {
    switch (character) {
    case 'K', 'k':
      if (character == 'k') {
        fmt.Println("Lowercase k at position", index)
        break
      }
      fmt.Println("Uppercase K at position", index)
    case 'y':
      fmt.Println("y at position", index)
    default:
      fmt.Println("character", string(character), "at position", index)
    }

  }
}
