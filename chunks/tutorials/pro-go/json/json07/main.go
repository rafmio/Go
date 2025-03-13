// Decoding JSON Data
package main

import (
	"encoding/json"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)

	vals := []interface{}{}

	decoder := json.NewDecoder(reader)

	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			Printfln("Error: %v", err.Error())
			break
		}
		vals = append(vals, decodedVal)
	}

	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}

	// Decoding Number Values
	Printfln("%s", "")

	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
				Printfln("Can't decode: %v, error: %v", ival, err)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}
}
