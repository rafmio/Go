// https://www.geeksforgeeks.org/methods-in-golang/
package main

import (
	"fmt"
)

type author struct {
	name   string
	branch string
}

func (a *author) show1(abranch string) {
	(*a).branch = abranch
}

func (a author) show2(abranch string) {
	a.name = "Gourav"
	fmt.Println("Author's name (Before):", a.name)
}

func main() {
	res := author{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("branch name (Before):", res.branch)

	res.show1("ECE")
	fmt.Println("branch name (After):", res.branch)

	(&res).show2("Mugu")
	// res.show2("Mugu")
	fmt.Println("Author's name (After):", res.name)
}
