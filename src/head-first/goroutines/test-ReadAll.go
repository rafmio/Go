package main

import (
  "fmt"
  "io"
  "log"
  "strings"
)

func main() {
  r := strings.NewReader("Go is a general-purpose language")
  fmt.Printf("Type of r: %T\n", r)

  b, err := io.ReadAll(r)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", b)

  sls := []int{1, 2, 3, 4, 5, 6, 7}
  fmt.Println("len of sls: ", len(sls))
}
