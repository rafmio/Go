// https://go101.org/article/type-embedding.html
package main

import "fmt"

type A struct {
	FieldX int
}

func (a A) MethodA() {
	fmt.Println("FieldX:", a.FieldX)
}

type B struct {
	*A
}

type C struct {
	B
}

func main() {
	var c = &C{B: B{A: &A{FieldX: 5}}}

	aa := c.B.A.FieldX
	bb := c.B.FieldX
	cc := c.A.FieldX
	dd := c.FieldX

	c.B.A.MethodA()
	c.B.MethodA()
	c.MethodA()
	fmt.Println(aa, bb, cc, dd)
}
