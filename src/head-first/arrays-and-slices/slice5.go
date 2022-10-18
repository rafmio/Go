package main

import "fmt"

func main() {
    s1 := []string{"s1", "s1"}
    s1 = append(s1, "s2", "s2")
    s1 = append(s1, "s3", "s3")
    s1 = append(s1, "s4", "s4")
    fmt.Println(s1)
    
    floatSlice  := make([]float64, 10)
    boolSlice   := make([]bool, 10)
    stringSlice := make([]string, 10)
    fmt.Println(floatSlice[5], boolSlice[5], stringSlice[5])
}
