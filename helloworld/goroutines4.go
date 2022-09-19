package main
import (
    "fmt"
    "time"
)

func Aname() {
    arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}
    
    for t1 := 0; t1 <= len(arr1) - 1; t1++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Printf("%s\n", arr1[t1])
    }
}

func Aid() {
    arr2 := [4]int{300, 400, 500, 1000}
    
    for t2 := 0; t2 <= len(arr2) - 1; t2++ {
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("%d\n", arr2[t2])
    }
}

func main() {
    fmt.Println("!...Main Go-routine Start...!")
    
    go Aname()
    
    go Aid()
    
    time.Sleep(3500 * time.Millisecond)
    fmt.Println("!...Main Go-routine End...!")
}

// Multiple goroutines
// https://www.geeksforgeeks.org/multiple-goroutines/
