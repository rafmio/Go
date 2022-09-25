package main

import "fmt"

type Numbers struct {
    num1 int
    num2 int
}

type NumbersInterface interface {
    Sum()   int 
    Mult()  int 
    Div()   float64  
    Sub()   int 
}

func (n Numbers) Sum() int {
    return n.num1 + n.num2
}

func (n Numbers) Mult() int {
    return n.num1 * n.num2
}

func (n Numbers) Div() float64 {
    return float64(n.num1) / float64(n.num2)
}

func (n Numbers) Sub() int {
    return n.num1 - n.num2
}
 

func main() {
    var i NumbersInterface
    numbers := Numbers{9, 8}
    i = numbers
    fmt.Printf("Sum: %d\n", i.Sum())
    fmt.Printf("Mul: %d\n", i.Mult())
    fmt.Printf("Div: %f\n", i.Div())
    fmt.Printf("Sub: %d\n", i.Sub())
    
    fmt.Println()
    fmt.Printf("Type of i: %T\nType of numbers: %T\n", i, numbers)
    
    fmt.Println()
    fmt.Println("i = ", i)
    fmt.Println("numbers = ", numbers)
    fmt.Println("numbers.Sum() : ", numbers.Sum())
}

// Interfaces 
// https://youtu.be/iMlJit_MfOA
