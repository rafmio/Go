package main

import (
  "fmt"
  "regexp"
)

func main() {
  description := "A boat for one person"

  match, err := regexp.MatchString("[A-z]oat", description)

  if (err == nil) {
    fmt.Println("Match:", match)
  } else {
    fmt.Println("Error:", err)
  }
  fmt.Println()

  pattern, compileErr := regexp.Compile("[A-z]oat")
  question := "Is that a goat?"
  preference := "I like oats"

  if (compileErr == nil) {
    fmt.Println("Description:", pattern.MatchString(description))
    fmt.Println("Question:", pattern.MatchString(question))
    fmt.Println("Preference:", pattern.MatchString(preference))
  } else {
    fmt.Println("Error:", compileErr)
  }
  fmt.Println()

}
