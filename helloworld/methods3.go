package main
import "fmt"

type Person struct {
    name string
    age int16 
}

func (p Person) Name() {
    fmt.Println(p.name, p.age)
}

func main() {
    p := Person{"Jack", 39}     // Create object
    p.Name()                    // Call method
}


// Methods
// https://golangdocs.com/methods-in-golang
