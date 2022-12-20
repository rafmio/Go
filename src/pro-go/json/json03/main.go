// Encoding Arrays and Slices
// Encoding Structs
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	encoder.Encode(m)

	fmt.Print(writer.String())

	// Kayak instance (type Product) from 'product.go' file
	var kayakWriter strings.Builder
	kayakEncoder := json.NewEncoder(&kayakWriter)
	kayakEncoder.Encode(Kayak)
	fmt.Print(kayakWriter.String())
	var kayakStr string = kayakWriter.String()
	fmt.Println(kayakStr)
	fmt.Println("---------------------------------------------")

	// Effect of promiotion in JSON in Encoding (embedded struct)
	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}

	var discWriter strings.Builder
	discEncoder := json.NewEncoder(&discWriter)
	discEncoder.Encode(dp)
	fmt.Print(discWriter.String())
}
