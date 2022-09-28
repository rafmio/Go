package main
import "fmt"

type Vertex struct {
    X int
    Y int
    Z int
    A int
}

func main() {
    v := Vertex{100, 200, 0, 0}
    fmt.Println(v)
    
    v.X = 104
    v.Y = 205
    fmt.Println(v)
    
    fmt.Printf("v.X = %d, v.Y = %d\n", v.X, v.Y)
    
    v.Z = v.X * v.Y
    fmt.Printf("v.Z = %d\n", v.Z)
    
    fmt.Printf("v.A = v.Z / v.Y: %d = %d / %d\n", v.Z / v.Y, v.Z, v.Y)
}


// Struct Fields
// https://go.dev/tour/moretypes/3
