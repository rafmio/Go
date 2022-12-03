package main

import (
  "fmt"
  "unicode"
  "strings"
)

func main() {
  product := "Kayakus"
  for _, char := range product {
    fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
  }
  fmt.Println()

  description := "A boat for one person"
  fmt.Println("Count:" ,strings.Count(description, "o"))
  fmt.Println("Index:", strings.Index(description, "o"))
  fmt.Println("Index:", strings.Index(description, "rs"))
  fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
  fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
  fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
  fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "o"))
  fmt.Println()

  // Inspecting strings with custom functions
  isLetterB := func (r rune) bool {
    return r == 'B' || r == 'b'
  }
  fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
}
