package main

import (
  "fmt"
  "strings"
)

const refString = "Mary_had a little_lamb"

func main() {
  words := strings.Split(refString, "_")
  for i, w := range words {
    fmt.Printf("Word %d is: %s\n", i, w)
  }
}
