// Decoding maps
package main

import (
	"encoding/json"
	"strings"
)

func main() {
	reader := strings.NewReader(`{"Kayak":279, "Lifejacket":49.85}`)

	m := map[string]float64{}

	decoder := json.NewDecoder(reader)

	err := decoder.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m, m)
		for k, v := range m {
			Printfln("Key: %v, value: %v", k, v)
		}
	}
}
