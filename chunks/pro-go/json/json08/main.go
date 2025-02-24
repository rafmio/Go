// Specifiying Types for Decoding
package main

import (
	"strings"
	"encoding/json"
)

func main() {
	reader := strings.NewReader(`true "Hello" 99.00 200`)

	var bval bool
	var sval string
	var fpval float64
	var ival int

	vals := []interface{} {&bval, &sval, &fpval, &ival}

	decoder := json.NewDecoder(reader)

	for i := 0; i < len(vals); i++ {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}

	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", ival, ival)
}
