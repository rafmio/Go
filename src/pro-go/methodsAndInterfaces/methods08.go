// Performing a type conversion
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275.00},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

func main() {
	products := ProductList(getProducts()) // conversion

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}
	fmt.Println()
	fmt.Println("----------------------------------------------")

	prodTest := getProducts() // without conversion
	for _, prod := range products {
		fmt.Println(prod)
	}
	fmt.Printf("Type of products: %T\n", products)
	fmt.Printf("Type of prodTest: %T\n", prodTest)
	fmt.Println()

	// will be an arror:
	// prodTest.calcCategoryTotals()

}

// You won't always be able to receive data with the type requiered to invoke
// a method defined for an alias. In these situation you can perform a type
// conversion
