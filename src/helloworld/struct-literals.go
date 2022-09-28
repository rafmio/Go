package main
import "fmt"

type Vertex struct {
    X, Y, Z int
}

var (
    v1 =  Vertex{100, 200, 300}
    v2 =  Vertex{X: 1}
    v3 =  Vertex{}
    p  = &Vertex{11, 12, 50}
)

func main() {
    fmt.Println(v1, p, v2, v3)
    fmt.Printf("Type of p is %T\n", p)
    fmt.Printf("Type of v1 is %T\n", v1)
    
    p.X = v1.X + v1.Y + v1.Z
    fmt.Printf("p.X = %d\n", p.X)
    
    v2 = Vertex{X: 111, Y: 300}
    fmt.Println("v2 = ", v2)
}

// Struct Literals
// https://go.dev/tour/moretypes/5
