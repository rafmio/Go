package main
import "fmt"

func f() func() int {
    i := 0
    return func() int {
        i += 10
        return i
    }
}

func main() {
    a := f()
    b := f()
    
    fmt.Println("a(): ", a())
    fmt.Println("b(): ", b())
    
    a()
    
    fmt.Println("a() :", a())
    fmt.Println("b() :", b())

    fmt.Println("a() :", a())
    fmt.Println("b() :", b())
    
    fmt.Println("a() :", a())
    fmt.Println("b() :", b())

    fmt.Println("a() :", a())
    fmt.Println("b() :", b())
    
    a()
    a()
    
    fmt.Println("a() :", a())
    fmt.Println("b() :", b())
    
}


// Function closure
// Data isolation of closure
// https://golangdocs.com/closures-in-golang
