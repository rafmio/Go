// https://habr.com/ru/companies/otus/articles/782414/
package main

import "fmt"

// Filter принимает слайс и функцию-предикт, возвращая новый слайс с элементами,
// удовлетворяющими условию
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	// Фильтруем слайс целых чисел
	ints := []int{1, 2, 3, 4, 5}
	even := Filter(ints, func(n int) bool { return n%2 == 0 })
	fmt.Println(even) // [2 4]

	// Фильтруем слайс строк
	str := []string{"apple", "banana", "cherry"}
	withA := Filter(str, func(s string) bool { return s[0] == 'a' })
	fmt.Println(withA) // ["apple"]
}
