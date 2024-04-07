// https://habr.com/ru/companies/otus/articles/782414/
package main

import (
	"fmt"
)

// PrintAny принимает слайс любых элементов и печатает их
func PrintAny(items []any) {
	for _, item := range items {
		fmt.Printf("%v ", item)
	}

	fmt.Println()
}

func main() {
	// Можем смешивать типы
	PrintAny([]any{"1", "2", "3"})
	PrintAny([]any{1, 2, 3})
	PrintAny([]any{1, "apple", true, 3.14})
}
