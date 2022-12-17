// https://pkg.go.dev/os
package main

import (
  "log"
  "os"
)

func main() {
  err := os.WriteFile("textFile.txt", []byte("Hello-mello, writed via WriteFile()\n"), 0666)
  if err != nil {
    log.Fatal(err)
  }
}
