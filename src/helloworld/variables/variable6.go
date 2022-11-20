package main
import "fmt"

func main() {
    myVar1, myVar2, myVar3 := 800, 34, 567
  
    fmt.Printf("The value of myvar1 is : %d\n", myVar1)
    fmt.Printf("The type of myvar1 is : %T\n", myVar1)
    
    fmt.Printf("\nThe value of myvar2 is : %d\n", myVar2)
    fmt.Printf("The type of myvar2 is : %T\n", myVar2)
    
    fmt.Printf("\nThe value of myvar3 is : %d\n", myVar3)
    fmt.Printf("The type of myvar3 is : %T\n", myVar3)
}
