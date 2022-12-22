// Appending to a file
package main

import (
  "os"
  "fmt"
)

func main() {
  f, err := os.Create("lines")
  if err != nil {
    fmt.Println(err.Error())
    f.Close()
  } else {
    f.Close()
  }

  f, err = os.OpenFile("lines", os.O_APPEND | os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println(err)
  }

  newLine := "File handling is easy"
  _, err = fmt.Fprintln(f, newLine)
  if err != nil {
    fmt.Println(err.Error())
    f.Close()
  }

  err = f.Close()
  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Println("file appended successfully")
}
