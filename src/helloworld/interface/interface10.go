// https://golangbyexample.com/interface-in-golang/
package main

import "fmt"

type animal interface {
  breathe()
  walk()
}

type lion struct {
  age int
}

func(l lion) breathe() {
  fmt.Println("Lion breathes")
}

func(l lion) walk() {
  fmt.Println("Lion walk")
}

func main() {
  var a animal
  a = lion{age: 10}
  a.breathe()
  a.walk()
  fmt.Printf("Type of a: %T\n", a)

  cat := lion{age: 10}
  cat.breathe()
  cat.walk()
}
