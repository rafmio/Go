// Decoding structs
package main

import (
	"encoding/json"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader(`
			{"Name":"Kayak","Category":"Watersports","Price":279, "Offer":"10"}
		`)

	decoder := json.NewDecoder(reader)

	for {
		var val DiscountedProduct
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v, Discount: %v",
				val.Name, val.Category, val.Price, val.Discount)
		}
	}
}
