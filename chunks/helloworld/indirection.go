package main

import "fmt"

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    fmt.Println("Inside Scale()")
    fmt.Println("f = ", f)
    fmt.Printf("Before Scale() acted: v.X = %f, v.Y = %f\n", v.X, v.Y)
    v.X = v.X * f
    v.Y = v.Y * f
    fmt.Printf("After  Scale() acted: v.X = %f, v.Y = %f\n", v.X, v.Y)
    fmt.Println("Ending of Scale()")
    fmt.Println("-------------------------------------------")
}

func ScaleFunc (v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(2)
    fmt.Printf("v = %v\n", v)
    ScaleFunc(&v, 10)
    fmt.Printf("v = %v\n", v)
}


// Methods and pointer indirections
// https://go.dev/tour/methods/6
