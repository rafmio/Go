// https://www.geeksforgeeks.org/methods-in-golang/
package main

import (
	"fmt"
)

type author struct {
	name   string
	branch string
}

func (a *author) show(abranch string) {
	(*a).branch = abranch
}

func main() {
	res := author{
		name:   "Sona",
		branch: "CSE",
	}

	p := &res

	fmt.Printf("Type of p: %T\n", p)
	fmt.Printf("Address of p: %p\n", p)
	fmt.Printf("Address of res: %p\n", &res)

	p.show("ECE")
	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch name (After): ", res.branch)
}
