package main

import (
    "fmt"
    "strings"
)

type myString string 

func (m myString) capitalize() myString {
    capStr := strings.ToUpper(string(m))
    
    return myString(capStr)
}

func main() {
    var m myString = "test"
    fmt.Println(m.capitalize())
    fmt.Println(m)
}

