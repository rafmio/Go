// Encoding JSON Data
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	var b bool = true
	var str string = "Hello-mello"
	var fval float64 = 99.99
	var ival int = 200
	var ptr *int = &ival

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	fmt.Println(encoder)
	fmt.Println()

	for _, val := range []interface{}{b, str, fval, ival, ptr} {
		encoder.Encode(val)
	}

	fmt.Printf("%#v\n", encoder)
	fmt.Println()

	fmt.Print(writer.String())
}
