// https://www.programiz.com/golang/empty-interface
package main

import "fmt"

func displayValue(i ...interface{}) {
	fmt.Println(i...)
	fmt.Println("----------------------------------------")
}

func main() {
	a := "Welcome, World!"
	b := 20
	c := 3.14
	d := false

	displayValue(a) // Output: Welcome, World!
	displayValue(a, b)
	displayValue(a, b, c, d) // Output: Welcome, World! 20 3.14 false
}
