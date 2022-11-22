// Defining and using methods
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category,
		"Price:", product.price)
}

func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275.00},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for _, p := range products {
		printDetails(p)
	}

	kayak := &Product{"Boat", "Watersports", 150.00}
	fmt.Println(kayak.name)
	fmt.Printf("Type of products: %T\n", products)
	fmt.Printf("Type of products[0]: %T\n", products[0])
	fmt.Printf("Type of products[0].name: %T\n", products[0].name)
}
