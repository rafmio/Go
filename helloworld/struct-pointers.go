package main
import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v
    p.X = 600
    p.Y = 400
    
    fmt.Println("p.X + p.Y = ", p.X + p.Y)
    
    (*p).X = 330
    (*p).Y = 440
    fmt.Println("p.X + p.Y = ", (*p).X + (*p).Y)
}
