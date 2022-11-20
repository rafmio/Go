package main

import "fmt"

func counter(start int) (func() int, func()) {
    ctr := func() int {
        return start
    }
    
    incr := func() {
        start++
    }
    
    return ctr, incr
}

func main() {
    ctr, incr := counter(100)
    ctr1, incr1 := counter(100)
    fmt.Println("counter - ", ctr())
    fmt.Println("counter1 = ", ctr1())
    
    incr()
    fmt.Println("counter - ", ctr())
    fmt.Println("counter1 - ", ctr1())
    
    incr1()
    incr1()
    fmt.Println("counter - ", ctr())
    fmt.Println("counter1 - ", ctr1())
    fmt.Println()
    
    ctr2, incr2 := counter(300)
    fmt.Println("ctr: ", ctr())
    fmt.Println("ctr1: ", ctr1())
    fmt.Println("ctr2: ", ctr2())
    incr2()
    fmt.Println("After incr2():")
    fmt.Println("ctr: ", ctr())
    fmt.Println("ctr1: ", ctr1())
    fmt.Println("ctr2: ", ctr2())
}


// https://kovardin.ru/articles/go/zamykaniya/
