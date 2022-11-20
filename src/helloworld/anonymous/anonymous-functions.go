package main
import "fmt"

// Passing an anonymous function as an argument into other 
// function
func GFG(i func(p, q string) string) {
    fmt.Println(i ("Mikle Jordan", "Chicago Bulls"))
}
    

func main() {
    func() {
        fmt.Println("Welcome, buddy!")
    } ()
    
    value := func () {
        fmt.Println("Metallica and AC/DC")
    }
    
    value()
    
    func(ele string) {
        fmt.Println(ele)
    } ("Joe Satriani")
    
    value2 := func(p, q string) string {
        return p + q + "Geeks"
    }
    GFG(value2)
}
