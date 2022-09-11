package main
import "fmt"

func main() {
    names := []string{
        "Kissy",
        "Missy",
        "Huggy",
        "Waggy",
        "Chikky",
        "Peakky",
    }
    fmt.Println(names)
    for i := 0; i < len(names); i++ {
        fmt.Println("name ", i, " : ", names[i])
    }
    
    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)
    
    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
}

// A slice goes not store any data, it just describes a 
// section of an underlying array
// Changing the elements of its underlying array
// Other slices that share the same underlying array will see 
// those changes
