package main

import "fmt"

type Contact struct {
  name string
  age int
}

func (x Contact) welcome() {
  fmt.Println(x.name)
  fmt.Println(x.age)
}

func main() {
  x := Contact{"James", 42}
  x.welcome()
}
