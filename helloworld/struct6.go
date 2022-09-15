package main
import "fmt"

type Author struct {
    name    string
    branch  string
    year    int
}

type HR struct {
    details Author
}

func main() {
    result := HR{details: Author{"Sona", "ECE", 2022}, }
    fmt.Println("Details of Author")
    fmt.Println(result)
}

// Nested structures
// https://www.geeksforgeeks.org/nested-structure-in-golang/
