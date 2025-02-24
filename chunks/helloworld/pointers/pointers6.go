package main
import "fmt"

func main() {
    // Example 1 ---------------------------------------
    var V     int = 100
    var pt1  *int = &V
    var pt2 **int = &pt1
    
    fmt.Println("V     = ", V)
    fmt.Println("&V    = ", &V)
    fmt.Println("pt1   = ", pt1)
    fmt.Println("*pt1  = ", &pt1)
    fmt.Println("pt2   = ", pt2)
    fmt.Println("*pt2  = ", *pt2)
    fmt.Println("**pt2 = ", **pt2)
    
    fmt.Println("--------------------------------------")
    // Example 2 ---------------------------------------
    var B     int = 100
    var pt3  *int = &B
    var pt4 **int = &pt3
    
    fmt.Printf("B = %d, pt3 = %v, pt4 = %v\n", B, pt3, pt4)
    *pt3 = 200
    fmt.Printf("B = %d, pt3 = %v, pt4 = %v\n", B, pt3, pt4)
    **pt4 = 300
    fmt.Printf("B = %d, pt3 = %v, pt4 = %v\n", B, pt3, pt4)
}



// Pointers to pointers (double pointers)
// https://www.geeksforgeeks.org/go-pointer-to-pointer-double-pointer/
