package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (g *Gallons) ToLiters() Liters {
    return Liters(*g * 3.785)
}

func (g *Gallons) ToMilleliters() Milliliters {
    return Milliliters(*g * 3_785.51)
}

func main() {
    milk := Gallons(2)
    fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk,
               milk.ToLiters())
    fmt.Printf("%0.3f gallons euqals %0.3f milliliters\n", milk, 
               milk.ToMilleliters())
}
