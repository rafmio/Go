package main
import "fmt"

func main() {
    i, j := 42, 2701
    
    p  := &i
    pp := &j
    
    fmt.Printf("*p = %v, *pp = %v\n", *p, *pp)
    
    *p  = 100
    *pp = 3000

    fmt.Printf("*p = %v, *pp = %v, i = %d, j = %d\n", *p, *pp, i, j)
    
    p = &j
    fmt.Println("*p = ", *p)
    *p = *p / 3 / 500
    fmt.Println("(After *p = *p / 3 / 500): *p = ", *p)
    
    fmt.Printf("i = %v, j = %v, *p = %v, *pp = %v\n", i, j, *p, *pp)
}


// Pointers 
// https://go.dev/tour/moretypes/1
