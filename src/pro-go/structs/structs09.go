// Defining anonymous struct types
package main

import "fmt"

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	type Item struct {
		name     string
		category string
		price    float64
	}

	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Stadium", category: "Soccer", price: 75000.00}

	writeName(prod)
	writeName(item)
}
