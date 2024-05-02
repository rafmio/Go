// https://golangdocs.com/recursion-in-golang
package main

import (
	"fmt"
)

func main() {
	var f func(int)

	f = func(v int) {
		if v == 5 {
			fmt.Println("Done")
			return
		} else {
			fmt.Println(v)
			f(v - 1)
		}
	}

	f(9)
}
