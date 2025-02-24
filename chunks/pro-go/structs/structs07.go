// Comparing struct values
package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}

	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p2:", p1 == p3)
}
