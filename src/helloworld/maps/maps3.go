package main
import "fmt"

func main() {
    products := map[string]float64 {
        "Java"      : 279  ,
        "JavaScript": 48.95,
    }
    
    fmt.Println("Map size: ", len(products))
    fmt.Println("Price: "   , products["Java"])
    fmt.Println("Price: "   , products["Hat"])
}

// Map
// Each map entry is specified using the key, a colon, the value, 
// and then a comma
