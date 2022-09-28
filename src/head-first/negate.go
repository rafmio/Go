package main

import "fmt"

func negate(myBoolean *bool) {
    *myBoolean = !*myBoolean
}


func main() {
    truth := true
    negate(&truth)
    fmt.Println("truth after negate() : ", truth)
    
    lies := false
    negate(&lies)
    fmt.Println("lies after negate() : ", lies)
}
