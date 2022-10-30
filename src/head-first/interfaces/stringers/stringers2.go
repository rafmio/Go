package main

import "fmt"

type Gallons float64 
func (g Gallons) String() string {
    return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64
func (l Liters) String() string {
    return fmt.Sprintf("%0.2f L", l)
}

type Milliliters float64 
func (m Milliliters) String() string {
    return fmt.Sprintf("%0.2f mL", m)
}

func main() {
    fmt.Println(Gallons(12.543435))
    fmt.Println(Liters(12.234032))
    fmt.Println(Milliliters(12.494323))
    
    a := Liters(10) + Liters(20)
    fmt.Println(a)
    fmt.Printf("Type of a: %T\n", a)
}
