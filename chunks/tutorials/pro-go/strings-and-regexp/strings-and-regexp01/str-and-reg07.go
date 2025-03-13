package main

import (
  "fmt"
  "strings"
)

func main() {
  text := "It was a boat. A small boat."

  elements := strings.Fields(text)
  fmt.Println("elements:", elements)
  fmt.Println("len(elements):", len(elements))
  joined := strings.Join(elements ,"--")
  fmt.Println("Joined:", joined)
  fmt.Println()

  var builder strings.Builder
  for _, sub := range strings.Fields(text) {
    if (sub == "small") {
      builder.WriteString("very ")
    }
    builder.WriteString(sub)
    builder.WriteRune(' ')
  }
  fmt.Println("Strings:", builder.String())
}
