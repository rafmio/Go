// Defining variadic parameters
package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Suppler:", supplier)
	}
}

func main() {
	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers("Lifejacket", "Sail Sage Co")

	// There is no output for the Soccer Ball product because nil slices have
	// zero length, so the for loop is never executed
	printSuppliers("Soccer Ball")
}

// The variadic parameter must be the last parameter defined by the function
