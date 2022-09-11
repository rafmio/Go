package main
import "fmt"

func main() {
    pow := make([]int, 10)
    fmt.Println(pow)
    for i := range pow {
        pow[i] = 1 << uint(i) // == 2 ** i
    }
    

    for _, value := range pow {
        fmt.Printf("%d ", value)
    } 
    
    fmt.Println("\n")
    
    for i := range pow {
        pow[i] = 0
    }
    
    for i := 0; i < len(pow); i++ {
        pow[i] = i * 2
    }
    fmt.Println(pow)
    
    fmt.Println("\n")
}


// You can skip the index or value by assigning to_
// for i, _ := range pow
// for _, value := range pow
// If you only want the index, you can omit the second variable
// for i := range pow
