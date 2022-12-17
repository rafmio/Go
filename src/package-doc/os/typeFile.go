// https://pkg.go.dev/os
package main

import (
  "os"
  "fmt"
)

func main() {
  var fln os.File = os.File{}
  fmt.Printf("Type of fln: %T\n", fln)
}
