package main

import (
  "fmt"
  "os"
  "encoding/json"
)

type Employee struct {
  FirstName string
  LastName string
  Post string
  Age int
  Residence string
}

func main() {
  jessie := &Employee{}

  file, err := os.Open("jsonJessie.json")
  if (err == nil) {
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&jessie)
    if (err != nil) {
      fmt.Println(err.Error())
    }
  } else {
    fmt.Println(err.Error())
  }

  defer file.Close()
  fmt.Println(jessie)
  fmt.Println(jessie.FirstName)
  fmt.Println(jessie.Age)
  fmt.Println(jessie.Residence)
}
