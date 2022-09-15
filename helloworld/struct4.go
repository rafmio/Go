package main
import "fmt"

type Car struct {
    Name, Model, Color string
    WeightInKg float64
}

func main() {
    c := Car{Name: "Ferrari", Model: "GTC4", Color: "Red", 
        WeightInKg: 1920}
    
    fmt.Println("Car Name: ", c.Name)
    fmt.Println("Car Model: ", c.Model)
    fmt.Println("Car Color: ", c.Color)
    fmt.Println("Weight in kg: ", c.WeightInKg)
    
    c.Color = "Black"
    fmt.Println("Car", c)
}

// https://www.geeksforgeeks.org/structures-in-golang/
