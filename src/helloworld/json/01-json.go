// https://golang.cafe/blog/golang-json-marshal-example.html
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a, err := json.Marshal(28192)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", a)            // Output: 28192
	fmt.Printf("Type of a: %T\n", a) // Output: []uint8

	var b int

}
