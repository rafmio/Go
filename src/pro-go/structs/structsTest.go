package main

import "fmt"

type ShipType struct {
  typeS string
  underwater bool
}

type Ship struct {
  name string
  *ShipType
}

func main() {
  submType := &ShipType{
    typeS: "submarine",
    underwater: true,
  }

  apl := &Ship{
    name: "APL",
    submType,
  }

  fmt.Println(apl)
}
