// https://pkg.go.dev/io
package main

import (
  "io"
  "log"
  "os"
  "strings"
  "fmt"
)

func main() {
  r := strings.NewReader("some io.Reader stream to be read")
  fmt.Printf("Type of r: %T\n", r)
  fmt.Println()

  if _, err := io.CopyN(os.Stdout, r, 4); err != nil {
    log.Fatal(err)
  }

  fmt.Println()
}
