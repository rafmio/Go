// https://www.geeksforgeeks.org/embedded-structures-in-golang/
package main

import (
	"fmt"
)

type Author struct {
	Name   string
	Branch string
	Year   int
}

type HR struct {
	Author
	id int
}

func main() {
	result := &HR{
		Author: Author{
			Name:   "Dhruv",
			Branch: "IT",
			Year:   2024,
		},
		id: 112233,
	}

	fmt.Println("Details fo Author")
	fmt.Println(result)
	fmt.Printf("Type of result: %T\n", result)
	fmt.Printf("result address: %p\n", result)
}
