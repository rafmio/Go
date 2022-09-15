package main
import "fmt"

type author struct {
    name    string
    branch  string
}

func (a *author) show_1(abranch string) {
    (*a).branch = abranch
}

func (a author) show_2() {
    a.name = "Gourav"
    fmt.Println("Author's name (Before): ", a.name)
}



// Methods can accept both pointer and value
// https://www.geeksforgeeks.org/methods-in-golang/
