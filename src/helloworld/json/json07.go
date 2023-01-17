// https://medium.com/rungo/working-with-json-in-go-7e3a37c5a07b
package main

import (
  "fmt"
  "encoding/json"
)

type Student map[string]interface{}

func main() {
  john := Student {
    "FirstName": "Jonh",
    "LastName": "Connor",
    "Age": 12,
    "HeightInMeters": 1.23,
    "IsMale": true,
  }

  johnJSON, err := json.Marshal(john)
  if (err != nil) {
    fmt.Println(err.Error())
  } else {
    fmt.Println(string(johnJSON))
  }
}
