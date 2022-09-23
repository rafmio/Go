package main
import "fmt"

func main() {
    TEST := 0
    cnt := func(a int) int {
        TEST = 2 + a
        return TEST
    }
    
    fmt.Println("The output of the closure func is: ", cnt(2))
    fmt.Println("The output of the closure func is: ", cnt(3))
    fmt.Println("The output of the closure func is: ", cnt(4))
}


// Function closure
// https://www.educba.com/golang-closure/
