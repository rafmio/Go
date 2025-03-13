package mypkg

import "fmt"

type MyInterface interface {
  MethodWithoutParemeters()
  MethodWithParameters(float64)
  MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParemeters() {
  fmt.Println("MethodWithParameters called")
}

func (m MyType) MethodWithParameters(f float64) {
  fmt.Println("MethodWithParameters", f)
}

func (m MyType) MethodWithReturnValue() string {
  return "Hi from MethodWithReturnValue"
}

func (m MyType) MethodNotItInterface() {
  fmt.Println("MethodNotItInterface called")
}
