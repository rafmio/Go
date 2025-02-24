package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"hello", "mello", "tosy", "bosy"}
	for i, x := range slices.Backward(s) {
		fmt.Println(i, x)
	}
}
