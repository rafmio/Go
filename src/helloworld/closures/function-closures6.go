package main
import "fmt"

func addition() func(int) int {
    add := 0
    return func(a int) int {
        add += a
        return add 
    }
}

func main() {
    positive, negative := addition(), addition()
    
    for k := 0; k < 11; k++ {
        fmt.Println(positive(k), "\t",  negative(-3 * k))
    }
}


// Function closure
// https://www.educba.com/golang-closure/
