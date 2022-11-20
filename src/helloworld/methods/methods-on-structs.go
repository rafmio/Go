package main
import "fmt"

type Human struct {
    name string
    age int
}

func (h Human) Describe() {
    fmt.Println(h.name, "is", h.age, "years old")
}

func main() {
    h := Human{"John", 23}
    h.Describe()
}


// Method on structs
// Methods can be defined by structs. It is really useful to do so
// since structs are the closest thing to a class in Go.
// They allow data encapsulation and with the methods described they
// will behave as a class does in an OOP language.
// https://golangdocs.com/methods-in-golang
