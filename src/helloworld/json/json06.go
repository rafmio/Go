package main

import (
  "os"
  "encoding/json"
  "fmt"
)

type Employee struct {
  FirstName string
  LastName string
  Post string
  Age int
  Residence string
}

type Suspects struct {
  Department string
  Agent string
  Suspects []Employee
}

func main() {
  jessie := &Employee {}
  walter := &Employee {}
  suspects := &Suspects {}

  var slsJesWal []*Employee = []*Employee{}

  file, err := os.Open("jsonJessie.json")
  if (err != nil) {
    fmt.Println("opening jessie", err.Error())
  } else {
    defer file.Close()
    decoder := json.NewDecoder(file) // jsonJessie.json
    err = decoder.Decode(&jessie)
    if (err != nil) {
      fmt.Println("decoding jessie:", err.Error())
    }
    slsJesWal = append(slsJesWal, jessie)
  }

  file, err = os.Open("jsonWalter.json")
  if (err != nil) {
    fmt.Println("opening walter", err.Error())
  } else {
    defer file.Close()
    decoder := json.NewDecoder(file) // jsonWalter.json
    err = decoder.Decode(&walter)
    if (err != nil) {
      fmt.Println("decoding walter:", err.Error())
    }
    slsJesWal = append(slsJesWal, walter)
  }

  file, err = os.Open("jsonSuspects.json")
  if (err != nil) {
    fmt.Println("opening suspects", err.Error())
  } else {
    defer file.Close()
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&suspects)
    if (err != nil) {
      fmt.Println("decoding suspects:", err.Error())
    }
  }

  fmt.Printf("%v\n", suspects)
  fmt.Println()

  for _, val := range slsJesWal {
    fmt.Println(val.FirstName)
    fmt.Println(val.LastName)
    fmt.Println(val.Age)
    fmt.Println()
  }
}
