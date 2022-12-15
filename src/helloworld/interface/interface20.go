// https://golangbyexample.com/interface-in-golang/
// Type switch
package main

import "fmt"

type animal interface {
  breathe()
  walk()
}

type lion struct {
  age int
}

func (l lion) breathe() {
  fmt.Println("Lion breathes")
}

func (l lion) walk() {
  fmt.Println("Lion walk")
}

type dog struct{
  age int
}

func (d dog) breathe() {
  fmt.Println("Dog breathes")
}

func (d dog) walk() {
  fmt.Println("Dog walk")
}

func main() {
  var a animal

  a = lion{age: 10}
  printA(a)

  var b animal
  b = dog{age: 7}
  printA(b)
}

func printA(a animal) {
  switch v := a.(type) {
  case lion:
    fmt.Printf("Type: %T\n", v)
  case dog:
    fmt.Printf("Type: %T\n", v)
  default:
    fmt.Printf("Unknown type: %T\n", v)
  }
}
