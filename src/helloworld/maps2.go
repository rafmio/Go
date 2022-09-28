package main
import "fmt"

func main() {
//     var names1 map[int]string    // name map has int keys and
                                 // string values
    var names2 = make(map[int]string)
    names2[0] = "John"
    names2[1] = "Jane"
    names2[2] = "Jerry"
    fmt.Println(names2)
    
    var names3 = map[int]string {
        0: "Jessy",
        1: "Wolter",
        // Last comma is a must
    }
    fmt.Println(names3)
    
    // You can insert keys in two ways
    // Indexed insertion
    var names4 = make(map[int]string)
    names4[0] = "Tukko"
    names4[1] = "Elladio"
    names4[2] = "Soul"
    fmt.Println(names4)
    fmt.Println(names4[0])
    fmt.Println(names4[1])
    fmt.Println(names4[2])
    
    // Check if a key exitst
    tukko, exists := names4[0]
    if(exists) {
        fmt.Printf("%s exitst", tukko)
        fmt.Printf("\n")
    }
    
    var names5 = make(map[int]string)
    names5[0] = "Freddy"
    names5[1] = "Shawn"
    names5[2] = "Jessy"
    names5[3] = "Guss"
    names5[4] = "Mike"
    names5[5] = "Fredd"
    fmt.Println("\n")
    for i := 0; i < len(names5); i++ {
        fmt.Println(names5[i])
    }
    fmt.Println("\n")
    for _, name := range names5 {
        fmt.Println(name)
    }
}


// Maps
// A map is a key-value pair storage container. It offers fast and 
// efficient lookups. And it doesn't allow duplicate keys while it
// can have duplicate values
//
// Maps can be initialized using map literals syntax.
