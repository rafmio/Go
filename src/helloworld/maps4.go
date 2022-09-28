package main
import "fmt"

func main() {
    keys := []string{}          // Slice
    
    carta := map[string][]int{
        "1": {1, 2},
        "2": {2, 10},
    }

    // Prints only keys
    for key := range carta {
        keys = append(keys, key)
        fmt.Println(key)
    }
    fmt.Println(keys)
    
    // Prints keys + values
    for key, value := range carta {
        fmt.Println(key, value)
    }
    
}


// Maps
