// https://pkg.go.dev/os
package main

import (
  "log"
  "os"
  "fmt"
)

func main() {
  data, err := os.ReadFile("testFile.txt")
  if err != nil {
    log.Fatal(err)
  }

  os.Stdout.Write(data)

  homeDir, err := os.UserHomeDir()
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Printf("homeDir: %v\n", homeDir)
  }
}
