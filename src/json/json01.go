package main

import (
  "fmt"
  "encoding/json"
)

type Message struct {
  Name string
  Body string
  Time int64
}

func main() {
  m := Message{"Alice", "Hello", 1294706395881547000}
  b, err := json.Marshal(m)
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Printf("%s\n", b)
    fmt.Printf("Type of m: %T\n", m)
    fmt.Printf("Type of b: %T\n", b)
  }

  fmt.Println()

  var decodeData Message
  err = json.Unmarshal(b, &decodeData)
  if err != nil {
    fmt.Println(err.Error())
  } else {
    fmt.Printf("%#v\n", decodeData)
  }
}
