package main

import "fmt"

type Car struct {
  name string
  color string
  year int
}

func main() {
  a := Car{"BMW", "red", 2018}
  p := &a
  fmt.Println(p.name)
}
