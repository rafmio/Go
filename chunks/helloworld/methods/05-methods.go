// https://www.geeksforgeeks.org/methods-in-golang/
package main

import (
	"fmt"
)

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {
	fmt.Println("Author's name:			", a.name)
	fmt.Println("Branch name:           ", a.branch)
	fmt.Println("Published articles:    ", a.particles)
	fmt.Println("Salary:				", a.salary)
}

func main() {
	res := author{
		name:      "Sona",
		branch:    "CSE",
		particles: 203,
		salary:    43_000,
	}

	res.show()
}
