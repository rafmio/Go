// https://go101.org/article/type-embedding.html
package main

import "fmt"

type X int

func (x X) M1() {
	fmt.Println(x)
}

func (x *X) M2() {
	fmt.Println(*x)
}

type T struct{ X }
type S struct{ *T }

func main() {
	var t = &T{X: 100}
	var s = S{T: t}
	var f = s.M1
	var g = s.M2
	s.X = 200
	f()
	g()
	s.T = &T{X: 300}
	f()
	g()
	fmt.Printf("Type of t: %T, type of s: %T\n", t, s)
}
