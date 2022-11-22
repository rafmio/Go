// Defining embedded fields
package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	type StockLevel struct {
		Product
		count int
	}

	stockItem := StockLevel{
		Product: Product{"Kayak", "Watersports", 275.00},
		count:   100,
	}

	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Category:", stockItem.Product.category)
	fmt.Println("Price:", stockItem.Product.price)
	fmt.Println("Count:", stockItem.count)
}

// If a field is defined without a name, it is known as an embedded field,
// and it is accessed using the name of its type
