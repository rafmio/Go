package main
import "fmt"

func main() { 
    var myVar1, myVar2, myVar3 = 2, "GFG", 4342.21
    
    fmt.Printf("myVar1: %d\n", myVar1)
    fmt.Printf("myVar1: %T\n", myVar1)
    fmt.Printf("myVar2: %s\n", myVar2)
    fmt.Printf("myVar2: %T\n", myVar2)
    fmt.Printf("myVar3: %f\n", myVar3)
    fmt.Printf("myVar3: %T\n", myVar3)
}
