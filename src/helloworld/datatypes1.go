package main
import "fmt"

func main() { 
    var X uint8 = 225
    fmt.Println(X, X-3) 

    var Y int16 = 32767
    fmt.Println(Y + 2, Y - 2)
    
    a := 20.45
    b := 34.89

    var n float32 = 11.11
    var m float32 = 44.44
    var o float32 = n + m
    fmt.Println("n + m : %f\n", o)
    
    c := b - a
    
    fmt.Printf("Result is: %f\n", c)
    fmt.Printf("The type of c is: %T\n", c)
    
}
