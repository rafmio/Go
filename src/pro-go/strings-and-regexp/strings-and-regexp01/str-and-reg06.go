package main

import (
  "fmt"
  "strings"
)

func main() {
  text := "It was a boat. A small boat."

  replace := strings.Replace(text, "boat", "canoe", 1)
  replaceAll := strings.ReplaceAll(text, "boat", "truck")

  fmt.Println("Replace:", replace)
  fmt.Println("replaceAll", replaceAll)
  fmt.Println()

  mapper := func(r rune) rune {
    if r == 'b' {
      return 'c'
    }
    return r
  }
  mapped := strings.Map(mapper, text)
  fmt.Println("Mapped:", mapped)
  fmt.Println()

  replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
  replaced := replacer.Replace(text)
  fmt.Println("Replaced:", replaced)
  fmt.Println(replacer)
  fmt.Printf("Type of replacer: %T\n", replacer)
  fmt.Printf("Type of replaced: %T\n", replaced)
  fmt.Println()
}
