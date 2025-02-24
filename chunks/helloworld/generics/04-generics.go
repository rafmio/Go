// https://habr.com/ru/companies/otus/articles/782414/
package main

import (
	"fmt"
)

type Number interface {
	int | float64 // may be int or float64
}

// UniversalAdd принимает два параметра любого типа, определенного в Number
func UniversalAdd[T Number](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(UniversalAdd[int](1, 2))

	fmt.Println(UniversalAdd[float64](1.1, 2.2))

}
