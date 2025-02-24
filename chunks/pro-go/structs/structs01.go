// Defining and using a struct
package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	kayak := Product{
		name:     "Kayak",
		category: "Waterports",
		price:    275,
	}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
}
