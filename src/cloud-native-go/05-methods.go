package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Square() {
	v.X *= v.X
	v.Y *= v.Y
}

func main() {
	vert := &Vertex{3, 4}
	fmt.Println(vert)
	fmt.Println("*vert:", *vert, "vert.X:", vert.X, "(*vert).X:", (*vert).X)

	vert.Square()
	fmt.Println(vert)
}
