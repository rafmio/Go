package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
  fmt.Println("Hi from", m)
}

func (m MyType) MethodWithParameters(number int, flag bool) {
  fmt.Println(m)
  fmt.Println(number)
  fmt.Println(flag)
}

func main() {
  value := MyType("a MyType value")
  value.sayHi()
  anotherValue := MyType("another value")
  anotherValue.sayHi()

  oneMoreValue := "one more value"
  value2 := MyType(oneMoreValue)
  value2.sayHi()

  value3 := MyType("MyType value")
  value3.MethodWithParameters(4, true)
}

// Необходимо привести значения к одному типу (MyType)
