package main

import "fmt"

type myStruct struct {
  myField int
}
func main() {
  var value myStruct
  value.myField = 3
  var pointer *myStruct = &value
  fmt.Println((*pointer).myField)
  fmt.Println(pointer.myField)
  fmt.Println(value.myField)

  fmt.Println()
  pointer.myField = 9
  fmt.Println((*pointer).myField)
  fmt.Println(pointer.myField)
  fmt.Println(value.myField)
}

// Если вместо (*pointer).myField разыменовывать как *pointer.myField,
// то Go подумает, что myField - это указатель
