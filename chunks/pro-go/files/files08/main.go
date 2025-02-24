// Create New Files
package main

import (
	"os"
	"encoding/json"
)

func main() {
	cheapProducts := []Product {}
	for _, p := range Products {
		if (p.Price < 100) {
			cheapProducts = append(cheapProducts, p)
		}
	}

	file, err := os.CreateTemp(".", "template-*.json")
	if (err == nil) {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
