package main

import "fmt"

type Gallons float64
type Liters float64
type Milliters float64

type CoffeePot string
func (c CoffeePot) String() string {
    return string(c) + " coffee pot" 
}

func main() {
    fmt.Println(Gallons(12.53285324))
    fmt.Println(Liters(12.03842430))
    fmt.Println(Milliters(12.032582342))
    
    coffeePot := CoffeePot("LuxBrew")
    fmt.Println(coffeePot)
    fmt.Println(coffeePot.String())
}
