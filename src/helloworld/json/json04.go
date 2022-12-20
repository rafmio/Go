// https://code.tutsplus.com/ru/tutorials/json-serialization-with-golang--cms-30209
package main

import (
  "fmt"
  "encoding/json"
)

func main() {
  // Serialize
  var x = 5
  bytes, err := json.Marshal(x)
  if err != nil {
    fmt.Println("Can't serialize", x)
  }

  fmt.Printf("%v => %v, '%v'\n", x, bytes, string(bytes))

  // Deserialize
  var r int
  err = json.Unmarshal(bytes, &r)
  if err != nil {
    fmt.Println("Can't deserialize", bytes)
    fmt.Println(err.Error())
  }

  fmt.Printf("%v => %v\n", bytes, r)
}
