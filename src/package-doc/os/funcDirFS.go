// https://pkg.go.dev/os
package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println(os.DirFS("/home"))  // added in go1.16
  
}
