package main

import (
  "fmt"
  "regexp"
)

func main() {
  pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat|[A-z]erson")
  description := "Kayak. A boat for one person."

  firstMatch := pattern.FindString(description)
  allMatches := pattern.FindAllString(description, -1)

  fmt.Println("First match:", firstMatch)

  for i, m := range allMatches {
    fmt.Println("Match:", i, "=", m)
  }
}
