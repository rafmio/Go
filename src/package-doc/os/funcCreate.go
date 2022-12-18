// https://pkg.go.dev/os
package main

import (
  "log"
  "os"
)

func main() {
  f, err := os.CreateTemp("", "example")
  if err != nil {
    log.Fatal(err)
  }

  defer os.Remove(f.Name())

  if _, err := f.Write([]byte("content")); err != nil {
    log.Fatal(err)
  }

  if err := f.Close(); err != nil {
    log.Fatal(err)
  }


}
