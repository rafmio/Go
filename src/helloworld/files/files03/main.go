// Writing strings line by line to a file
package main

import (
  "fmt"
  "os"
)

func main() {
  f, err := os.Create("lines")
  if (err != nil) {
    fmt.Println(err.Error())
  }

  d := []string{"Welcome to the world of Go!", "Go is a compiled language.", "It is easy to learn Go."}

  for _, v := range d {
    fmt.Fprintln(f, v)
  }

  fmt.Println()
  for i, v := range d {
    fmt.Println(i, ":", v)
  }
  fmt.Println()

  err = f.Close()
  if (err != nil) {
    fmt.Println(err.Error())
  }

  fmt.Println("file written successfully")
}
