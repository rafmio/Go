package main

import "fmt"

func main() {
	var price int = 100
	var currency string = "dollars"
	fmt.Println("Price is", price, currency)

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, currency)

	var total float64 = float64(price) + tax
	fmt.Println("Total costs is", total, currency)

	var availableFuns int = 120
	fmt.Println(availableFuns, currency, "available")
	fmt.Println("Within budget?", float64(availableFuns) <= total)
}
