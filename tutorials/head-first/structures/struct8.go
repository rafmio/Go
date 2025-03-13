package main

import "fmt"

type part struct {
  description string
  count int 
}

func minimumOrder(description string) part {
  var p part
  p.description = description
  p.count = 100
  return p
}

func main() {
  p := minimumOrder("Hex bolts")
  fmt.Println(p.description, p.count)
}
