// https://pkg.go.dev/strings
package main

import (
  "fmt"
  "strings"
)

func main() {
  var b strings.Builder
  for i := 3; i >= 1; i-- {
    fmt.Fprintf(&b, "%d...", i)
  }

  b.WriteString("ignition")
  fmt.Println(b.String())

  fmt.Println()
  var aA strings.Builder
  fmt.Fprintf(&aA, "%s", "Hehehe it's coool")
  fmt.Println(aA)
}
