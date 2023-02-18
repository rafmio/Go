package main

import (
  // "io"
  "fmt"
  "os"
)

func main() {
  if err := os.WriteFile("original.txt", []byte("String to write to file\n"), 0666); err != nil {
    fmt.Println("writing to file: ", err.Error())
  }
}
