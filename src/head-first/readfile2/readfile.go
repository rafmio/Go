package main

import (
    "fmt"
    "datafile2"
    "log"
)

func main() {
    numbers, err := datafile2.GetFloats("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    var sum float64 = 0
    for _, number := range numbers {
        sum += number 
    }
    
    sampleCount := float64(len(numbers))
    fmt.Printf("Average: %0.2f\n", sum / sampleCount)
}
