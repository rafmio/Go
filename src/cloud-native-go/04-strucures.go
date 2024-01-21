package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func main() {
	var v Vertex
	fmt.Println(v)

	v = Vertex{}
	fmt.Println(v)

	v = Vertex{1.0, 2.0}
	fmt.Println(v)

	v = Vertex{Y: 2.5}
	fmt.Println(v)

	v.X = 300
	fmt.Println(v)

	var b *Vertex = &Vertex{11, 33}
	fmt.Println(b)
	fmt.Println(*b)
}
