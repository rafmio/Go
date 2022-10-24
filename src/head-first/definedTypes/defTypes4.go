package main

import "fmt"

type MyType string

func (m MyType) sayHi () {
  fmt.Println("Hi")
}

func main() {
  value := MyType("a MyType value")
  fmt.Println("value: ", value)
  value.sayHi()
  anotherValue := MyType("another value")
  fmt.Println("anotherValue: ", anotherValue)
  anotherValue.sayHi()
}
