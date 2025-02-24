package main

import "fmt"

func main() {
  a := "Hello"
  b := 'A'
  c := `D`
  d := 123
  fmt.Println(a, b, c, d)
  fmt.Printf("%v %v %v %v\n", a, b, c, d)
  fmt.Printf("%#v %#v %#v %#v\n", a, b, c, d)
  fmt.Printf("%c\n", b)
}
