package main

import "fmt"

func main() {
    var testVar interface{} = 42.56
    
    switch t := testVar.(type) {
        case string:
            fmt.Println("Var is of type string")
        case int: 
            fmt.Println("Var is of type int")
        case float32:
            fmt.Println("Var is of type float32")
        default:
            fmt.Printf("Var type is of %T and unknown\n", t)
    }
    
}

// https://blog.logrocket.com/type-assertions-vs-type-conversions-go/
