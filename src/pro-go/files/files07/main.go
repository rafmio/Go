// Writing JSON Data to a File
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

	file, err := os.OpenFile("cheap.json", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
	if (err == nil) {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
