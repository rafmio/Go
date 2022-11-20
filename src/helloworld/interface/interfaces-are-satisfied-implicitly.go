package main

import "fmt"

type Intfc interface {
  M()
  MM()
}

type Strct struct {
  msg   string
  num   int 
}

func (t *Strct) M() {
  fmt.Println("t.msg: ", t.msg)
}

func (t *Strct) MM() {
  t.num = t.num * 100
  fmt.Println("t.num: ", t.num)
}

func main() {
  var i Intfc = &Strct{"Tosi-Bosi", 24}
  i.M()
  i.MM()
}

// Interfaces are implemented implicitly
// https://go.dev/tour/methods/10
