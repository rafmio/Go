package main
import "fmt"

type Vertex struct {
    X int
    Y int
    Z int
}

func main () {
    v := Vertex{10, 20, 0}
    p := &v
    p.X = 1e9
    fmt.Println(v)
    
    p.Z = p.X / p.Y / 10_000
    fmt.Printf("p.Z = %d\n", p.Z)
}


// Pointers to struct
// https://go.dev/tour/moretypes/4
