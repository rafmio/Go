package main
import "fmt"

func main() { 
    Element := struct { 
        name        string
        branch      string
        language    string
        Particles   int
    }{
        name:       "Packachu",
        branch:     "ECE",
        language:   "Golang",
        Particles:  498,
    }
    
    fmt.Println(Element)
    
} 


// Anonymous structures
//https://www.geeksforgeeks.org/anonymous-structure-and-field-in-golang/
