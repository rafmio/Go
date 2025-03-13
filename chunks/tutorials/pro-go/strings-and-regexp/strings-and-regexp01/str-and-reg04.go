package main

import (
  "fmt"
  "strings"
)

func main() {
  description := "A boat for one person"

  splits := strings.SplitN(description, " ", 3)
  for _, x := range splits {
    fmt.Println("Splits >>" + x + "<<")
  }
  fmt.Println()

  description2 := "This  is  a  double  spaced"
  splits2 := strings.SplitN(description2, " ", 3)
  for _, x := range splits2 {
    fmt.Println("Split >>" + x + "<<")
  }
  fmt.Println()

  splits3 := strings.Fields(description2)
  for _, x := range splits3 {
    fmt.Println("Field >>" + x + "<<")
  }
  fmt.Println()

  splitter := func(r rune) bool {
    return r == ' '
  }
  splits4 := strings.FieldsFunc(description2, splitter)
  for _, x := range splits4 {
    fmt.Println("Field >>" + x + "<<")
  }
}
