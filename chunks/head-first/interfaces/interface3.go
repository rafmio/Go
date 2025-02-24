package main

import "fmt"

type Switch string
func (s *Switch) toggle() {
  if *s == "on" {
    *s = "off"
  } else {
    *s = "on"
  }
  fmt.Println(*s)
}

type Toggleable interface {
  toggle()
}

func main() {
  s := Switch("off")
  fmt.Printf("Type of s: %T\n", s)
  var t Toggleable = &s
  fmt.Printf("Type of t: %T\n", t)
  t.toggle()
  t.toggle()
  t.toggle()
}
