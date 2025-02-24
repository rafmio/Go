package main

import (
	"fmt"
)

// Passing an anonymous function as an argument into other func
func GFG(i func(p, q string) string) {
	fmt.Println(i("Mikle Jordan", "Chicago Bulls"))
}

func main() {
	func() {
		fmt.Println("Welcome buddy!")
	}()

	value := func() {
		fmt.Println("Metallica and AC/DC")
	}

	value()

	func(ele string) {
		fmt.Println(ele)
	}("Joe Satriani")

	value2 := func(p, q string) string {
		return p + q + "Geeks"
	}

	fmt.Printf("type of value2 is: %T\n", value2)

	GFG(value2)
}
