package main

import "fmt"

func main() {
	var subscriber1 struct {
		name   string
		rate   float64
		active bool
	}

	subscriber1.name = "Aman Singh"
	subscriber1.rate = 5.55
	fmt.Println("Name: ", subscriber1.name, "Rate: ", subscriber1.rate)
	fmt.Printf("%#v\n", subscriber1)

	var subscriber2 struct {
		name   string
		rate   float64
		active bool
	}
	subscriber2.name = "Beth Ryan"
	subscriber2.rate = 56.11
	subscriber2.active = true

	fmt.Println()
	fmt.Println("Name: ", subscriber2.name)
	fmt.Printf("%#v\n", subscriber2)
}
