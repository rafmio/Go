// https://golangbot.com/write-files/
package main

import (
  "fmt"
  "os"
)

func main() {
  f, err := os.Create("text.txt")
  if (err != nil) {
    fmt.Println(err.Error())
  }

  w, err := f.WriteString("Hello Kissy-Missy\n")
  if (err != nil) {
    fmt.Println(err.Error())
  }

  fmt.Println(w, "bytes written successfully")
  err = f.Close()
  if (err != nil) {
    fmt.Println(err.Error())
  }
}
