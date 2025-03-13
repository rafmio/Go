// Using named results
package main

import "fmt"

func calcTax(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func calcTotalPrice(products map[string]float64, minSpend float64) (
	total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

// total = 10 + 55 + 39 + 9.5 + 48.95
// tax = 55 + 39
func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Tent":       195,
		"Hat":        9.5,
		"Lifejacket": 48.95,
	}

	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)
}
