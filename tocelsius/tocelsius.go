package main

import (
    "fmt"
    "keyboard"
    "log"
)

func main() {
    fmt.Print("Enter a temperature in Fahrenheit: ")
    fahr, err := keyboard.GetFloat()
    if err != nil {
        log.Fatal(err)
    }
    
    celsius := (fahr - 32) * 5 / 9
    fmt.Printf("%0.2f degrees Celsius\n", celsius)
    
    fmt.Println()
    fmt.Println("Finally!")
}
