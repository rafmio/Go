// https://habr.com/ru/companies/otus/articles/782414/
package main

import "fmt"

// Distinct позволяет убедиться, что все элементы в слайсе уникальны
func Distinct[T comparable](list []T) []T {
	unique := make(map[T]bool)
	var result []T
	for _, item := range list {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}

	return result
}

func main() {
	fmt.Println(Distinct([]int{1, 2, 3, 4, 5, 6, 7, 11, 1}))
	fmt.Println(Distinct([]string{"Huggy", "Wuggy", "Kissy", "Missy", "Huggy"}))
}
