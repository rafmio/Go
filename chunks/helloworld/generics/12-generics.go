// https://go.dev/tour/generics/1
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here
		if v == x {
			return i
		}
	}

	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz", "hello"}
	fmt.Println(Index(ss, "hello"))

	sf := []float64{1.3, 5.5, 300.0, 500.01}
	fmt.Println(Index(sf, 3.14))
}
