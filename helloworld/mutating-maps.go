package main
import "fmt"

func main() {
    m := make(map[string]int)
    m["Answer"] = 42
    fmt.Println("The value: ", m["Answer"])
    
    m["Answer"] = 48
    fmt.Println("The value: ", m["Answer"])
    
    delete(m, "Answer")
    fmt.Println("The value: ", m["Answer"])
    
    v, ok := m["Answer"]
    fmt.Println("The value: ", v, "Present?", ok)
}


// Mutating maps
// https://go.dev/tour/moretypes/22
// Insert or update an element in map 'm':
// m[key] = elem 
// Retrieve an element: 
// elem = m[key]
// Delete an element: 
// delete(m, key)
// Test that a key is present with a two-value assignment:
// elem, ok = m[key]
