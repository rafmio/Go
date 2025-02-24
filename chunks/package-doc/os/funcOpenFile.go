// https://pkg.go.dev/os
package main

import (
  "log"
  "os"
)

func main() {
  f, err := os.OpenFile("notes.txt", os.O_RDWR | os.O_CREATE, 0666)
  if err != nil {
    log.Fatal(err)
  }


  if err := f.Close(); err != nil {
    log.Fatal(err)
  }
}
