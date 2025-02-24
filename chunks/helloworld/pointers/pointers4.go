package main
import "fmt"

func ptf(a *int) {
    *a = 748
}

func ptf2(a *int) {
    *a = 748
}

func main() {
    // Example 1------------------------------------------
    var x1 = 100
    fmt.Println("x1 before passing to ptf: ", x1)
    var pa *int = &x1
    
    ptf(pa)
    fmt.Println("x1 after passing to ptf:  ", x1)

    fmt.Println("----------------------------------------")
    // Example 2------------------------------------------
    var x2 = 100
    fmt.Println("x2 before passing to ptf: ", x2)
    ptf(&x2)
    fmt.Println("x2 after passing to ptf:  ", x2)
}

// Pointers to a funcions
// https://www.geeksforgeeks.org/pointers-to-a-function-in-go/
