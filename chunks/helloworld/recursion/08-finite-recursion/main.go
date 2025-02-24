// https://golangdocs.com/recursion-in-golang
// finite recursion
package main

import (
	"fmt"
)

func recurFact(v int) int {
	if v == 0 {
		return 1
	} else {
		result := v * recurFact(v-1)
		fmt.Printf("v: %d, result: %d\n", v, result)
		return result
	}
}

func main() {
	n := recurFact(10)
	fmt.Println(n)
}
