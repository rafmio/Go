// https://pkg.go.dev/strings
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Compare("a", "b"))
  fmt.Println(strings.Compare("a", "a"))
  fmt.Println(strings.Compare("b", "a"))
  fmt.Println(strings.Compare("Kissy-Missy", "Kissy-Missy"))
  fmt.Println(strings.Compare("Kissy", "Missy"))
}
