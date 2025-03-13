  package main

  import (
    "fmt"
    "mypkg"
  )

  func main() {
    var value mypkg.MyInterface
    value = mypkg.MyType(5)
    fmt.Printf("Value of 'value' var is: %v\n", value)
    value.MethodWithoutParemeters()
    value.MethodWithParameters(127.3)
    fmt.Println(value.MethodWithReturnValue())

    fmt.Println()
    fmt.Printf("Type of value var: %T\n", value)
  }
