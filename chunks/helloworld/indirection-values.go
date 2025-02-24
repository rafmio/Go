package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) AbsMet() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func AbsFunc(v Vertex) float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
    v := Vertex{5, 7}
    fmt.Println(v.AbsMet())
    fmt.Println(AbsFunc(v))
    
    p := &Vertex{10, 20}
    fmt.Println(p.AbsMet())
    fmt.Println(AbsFunc(*p))
    
    fmt.Printf("Type of v: %T, type of p: %T\n", v, p)
}

// Methods and pointers indirection (2)
// https://go.dev/tour/methods/7
