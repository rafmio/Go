// https://pkg.go.dev/strings
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
  fmt.Println()

  str := "?String for trimming,"

  fmt.Print(strings.Trim(str, "?,"))
  fmt.Println()
}
