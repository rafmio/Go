// https://www.geeksforgeeks.org/methods-in-golang/
package main

import (
	"fmt"
)

type data int

func (d1 data) multyply(d2 data) data {
	return d1 * d2
}

func main() {
	value1 := data(23)
	value2 := data(20)
	res := value1.multyply(value2)

	fmt.Println("Final result: ", res)
}
