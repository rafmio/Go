// https://pkg.go.dev/strings
package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar baz   "))

  str := "Ехал Грека через реку видит Грека в реке рак"
  splt := strings.Fields(str)
  fmt.Println(splt)
  fmt.Printf("Type of splt: %T\n", splt)
}
