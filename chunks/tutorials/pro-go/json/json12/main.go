// Decoding structs
package main

import (
	"strings"
	"encoding/json"
	"io"
)

func main() {
	reader := strings.NewReader(`
			{"Name":"Kayak","Category":"Watersports","Price":279}
			{"Name":"Lifejacket","Category":"Watersports"}
			{"name":"Canoe","category":"Watersports","price":100,"isStock":true}
		`)

	decoder := json.NewDecoder(reader)

	slsProd := []Product{} // my

	for {
		var val Product
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v",
			val.Name, val.Category, val.Price)
			slsProd = append(slsProd, val) // my
		}
	}

	// my :
	Printfln("%s", "")
	Printfln("Type of slsProd: %T", slsProd)
	for _, v := range slsProd {
		Printfln("%v, %v, %v", v.Name, v.Category, v.Price)
	}
}
