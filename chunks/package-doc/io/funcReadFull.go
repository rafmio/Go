// https://pkg.go.dev/io
// func ReadFull(r Reader, buf []byte) (n int, err error)
package main

import (
  "fmt"
  "io"
  "log"
  "strings"
  "os"
)

func main() {
  r := strings.NewReader("some io.Reader stream to be read\n")
  fmt.Printf("Type of r: %T\n", r)
  fmt.Println()

  buf := make([]byte, 4)
  if _, err := io.ReadFull(r, buf); err != nil {
    log.Fatal(err)
  }
  io.Copy(os.Stdout, r)
  fmt.Println(string(buf))
  fmt.Println("len(buf): ", len(buf))

  // minimal read size bigger than io.Reader stream
  longBuf := make([]byte, 64)
  if _, err := io.ReadFull(r, longBuf); err != nil {
    fmt.Println("error:", err)
  }

}
