package main
import "fmt"

type author struct {
    name        string
    branch      string
}

func (a *author) show(abranch string) {
    (*a).branch = abranch
}

func main() {
    res := author {
        name:   "Sona",
        branch: "CSE",
    }
    
    p := &res
    
    p.show("ECE")
    fmt.Println("Author's name: ", res.name)
    fmt.Println("Branch name (After): ", res.branch)
}



// Methods with Pointer Receiver
// https://www.geeksforgeeks.org/methods-in-golang/
