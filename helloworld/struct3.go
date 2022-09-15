package main
import "fmt"

type Address struct {
    Name    string
    City    string
    Pincode int
}

func main() {
    var a Address
    fmt.Println(a)
    
    a1 := Address{"Akshay", "Dehradun", 3623572}
    fmt.Println("Address 1: ", a1)

    a2 := Address{Name: "Anikaa", City: "Ballia", Pincode: 277001}
    fmt.Println("Address 2: ", a2)
    
    a3 := Address{Name: "Delhi"}
    fmt.Println("Address 3: ", a3)
    
    var a4 Address = Address {Name:"Raf", City: "Kazan", Pincode: 420330}
    fmt.Println("Address 4: ", a4)
}

// Structures
// https://www.geeksforgeeks.org/structures-in-golang/
