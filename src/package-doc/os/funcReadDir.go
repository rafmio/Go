// https://pkg.go.dev/os
package main

import (
  "fmt"
  "log"
  "os"
)

func main() {
  files, err := os.ReadDir(".")
  if err != nil {
    log.Fatal(err)
  }

  filesRoot, err := os.ReadDir("/home/raf-pc/Go/src")
  if err != nil {
    log.Fatal(err)
  }

  for _, file := range files {
    fmt.Println(file.Name())
  }

  fmt.Println()

  for _, file := range filesRoot {
    fmt.Println(file.Name())
  }

}
