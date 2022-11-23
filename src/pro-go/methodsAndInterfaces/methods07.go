// Defining methods for type aliases
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

// Код работает как со звездочкой *, так и без нее
type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func main() {
	products := ProductList{
		{"Kayak", "Watersports", 275.00},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
		{"Whistle", "Soccer", 1.30},
	}

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}

	fmt.Println()
	returnedMap := products.calcCategoryTotals()
	fmt.Println(returnedMap)
	fmt.Printf("Type if returnedMap: %T\n", returnedMap)
}
