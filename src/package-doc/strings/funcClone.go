// https://pkg.go.dev/strings
// There is no Clone() func at go 1.15 version
package main

import (
  "fmt"
  "strings"
)

func main() {
  str := "A string for clone()"
  str2 := strings.Clone(str)
  fmt.Println(str)
  fmt.Println(str2)
}
