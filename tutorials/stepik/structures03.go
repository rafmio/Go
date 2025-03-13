package main

import "fmt"

type T struct {
  val int
}

func (p *T) a() {
  p.val += 1
}

func (p T) b() {
  p.val += 2
}

func main() {
  x := T{5}
  x.a()
  x.b()
  fmt.Println(x.val)
}
