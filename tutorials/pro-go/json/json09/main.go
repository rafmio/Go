// Decoding Arrays
package main

import (
	"encoding/json"
	"strings"
	"io"
)

func main() {
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)

	vals := []interface{} {}

	decoder := json.NewDecoder(reader)

	var iterations int = 0

	for {
		var decodedVal interface {}
		err := decoder.Decode(&decodedVal)
		if (err != nil) {
			if (err != io.EOF) {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
		iterations++
	}

	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}
	Printfln("interations: %v", iterations)
}
