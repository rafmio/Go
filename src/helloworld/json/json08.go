// https://drstearns.github.io/tutorials/gojson/
package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Rectangle struct {
  Top int
  Left int
  Width int
  Height int
}

func main() {
  r := &Rectangle{10, 20, 30, 40}

  enc := json.NewEncoder(os.Stdout)
  if err := enc.Encode(r); err != nil {
    fmt.Printf("error encoding struct into JSON: %v\n", err)
  }
}
