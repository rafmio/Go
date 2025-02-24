package main

import (
	"fmt"
)

func testcount(x int) int {
	// base case
	if x == 11 {
		return 0
	}

	// recursive case? рекурсивное отношение
	fmt.Printf("%d ", x)
	return testcount(x + 1)
}

func main() {
	testcount(1)
	fmt.Println()
}
