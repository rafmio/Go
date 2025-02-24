// Using an initializaion
package main

import "fmt"

func calcTax(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
		"Hat":        6.18,
		"Tent":       196.50,
	}

	for product, price := range products {
		if taxAmount, taxDue := calcTax(price); taxDue {
			fmt.Println("Product:", product, "Tax:", taxAmount)
		} else {
			fmt.Println("Product:", product, "No tax value")
		}
	}
}
