// Undertanding pointers to values
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func calcTax(product *Product) {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
}

func main() {
	kayak := &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275.00,
	}

	fmt.Printf("Type of kayak: %T\n", kayak)
	calcTax(kayak)

	fmt.Println("Name:", kayak.name, "Category:", kayak.category,
		"Price:", kayak.price)

}

// There is no need to assing a struct value to a variable before creating
// a pointer, and the address operator cab used directly with the literal
// struct syntax
