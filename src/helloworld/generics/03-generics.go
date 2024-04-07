// https://habr.com/ru/companies/otus/articles/782414/
package main

import (
	"fmt"
)

func Merge[T any, U any](first T, second U) {
	fmt.Println(first, second)
}

func main() {
	// явное указание типов необходимо, т.к. Go не может однозначно
	// вывести их
	Merge[int, string](42, "The answer is")
	Merge[string, int]("The answer is", 42)
}
