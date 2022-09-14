package main
import (
    "fmt"
    "strings"
)

func joinstr(elements ...string) string {
    return strings.Join(elements, "-")
}

func main() {
    fmt.Println(joinstr())
    fmt.Println(joinstr("Geek", "GfG", "Newby"))
    fmt.Println(joinstr("C", "h", "e", "e", "e", "e", "e", "s", "e"))
    
    // Pass a slice in variadic functioin
    elements := []string{"veni", "vidi", "viki"}
    fmt.Println(joinstr(elements...))
}
// variadic functions
//https://www.geeksforgeeks.org/variadic-functions-in-go/
