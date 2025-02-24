package main
import "fmt"

func main() {
    val1 := 2345
    val2 := 567
    
    var p1 *int
    p1 = &val1
    p2 := &val2
    p3 := &val1
    
    fmt.Println("&p1 == &p2: %v", &p1 == &p2)
    fmt.Println(" p1 ==  p2: %v", p1 == p2)
    fmt.Println(" p1 ==  p3: %v", p1 == p3)
    fmt.Println(" p2 ==  p3: %v", p2 == p3)
    fmt.Println("&p3 == &p1: %v", &p3 == &p1)
}
