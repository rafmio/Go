// Assigning a value to an anonymous struct
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod.name,
		ProductPrice: prod.price,
	})

	fmt.Println(builder.String())
}
