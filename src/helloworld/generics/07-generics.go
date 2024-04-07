// https://habr.com/ru/companies/otus/articles/782414/
package main

import (
	"fmt"
)

// Swap меняет местами значения двух переменных любого типа
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func main() {
	x := 1
	y := 2
	fmt.Printf("origin values: x = %v, y = %v\n", x, y)

	Swap(&x, &y)
	fmt.Printf("swapped values: x = %v, y = %v\n", x, y)

	s1 := "Huggy"
	s2 := "Wuggy"
	fmt.Printf("origin values: x = %s, y = %s\n", s1, s2)

	Swap(&s1, &s2)
	fmt.Printf("swapped values: x = %s, y = %s\n", s1, s2)

}
