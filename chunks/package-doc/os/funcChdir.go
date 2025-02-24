// https://pkg.go.dev/os
package main

import (
  // "fmt"
  "strings"
  "os"
  "io"
  "fmt"
)

func main() {
  text := strings.NewReader("Changing directory...\n")
  io.Copy(os.Stdout, text)

  path := "/home/"  // I can't figure out how it works
  if err := os.Chdir(path); err != nil {
    errTxt := strings.NewReader("Error occuring during changing directory\n")
    io.Copy(os.Stdout, errTxt)
    fmt.Println(err)
  }
}
