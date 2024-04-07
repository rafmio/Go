// https://habr.com/ru/companies/otus/articles/782414/
package main

import (
	"fmt"
)

// Generic функция, которая работает с любым типом
func Print[T any](s T) {
	fmt.Println(s)
}

func main() {
	// Go выводит тип параметра T как int
	Print(123)

	Print("Hello, Generics!")
}
