package main
import "fmt"

type author struct {
    name    string
    branch  string
}

// Method show_1 -----------------------------------
func (a *author) show_1(abranch string) {
    (*a).branch = abranch
}

func (a author) show_2() {
    a.name = "Gourav"
    fmt.Println("Author's name (Before): ", a.name)
}

func main() {
    res := author {
        name:   "Sona",
        branch: "CSE",
    }
    
    fmt.Println("branch name (Before): ", res.branch)
    
    res.show_1("ECE")
    fmt.Println("branch name (After): ", res.branch)
    
    (&res).show_2()
    fmt.Println("Author's name (After): ", res.name)
}


// Methods can accept both pointer and value
// https://www.geeksforgeeks.org/methods-in-golang/
