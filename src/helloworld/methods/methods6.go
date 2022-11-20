package main
import "fmt"

type data int                               // Type definition

func (d1 data) multiply (d2 data) data {    // Defining a method    
    return d1 * d2                          // with non-struct type
}                                           // receiver

func main() {
    value1 := data(23)
    value2 := data(20)
    res    := value1.multiply(value2)
    
    fmt.Println("Final result: ", res)
}

// Method with Non-Struct Type Receiver
// https://www.geeksforgeeks.org/methods-in-golang/
