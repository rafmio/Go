// https://habr.com/ru/companies/otus/articles/782414/
package main

import (
	"fmt"
)

// Stack представляет собой универсальный стек
type Stack[T any] struct {
	elements []T
}

// Push добавляет элемент в стек
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop удаляет и возвращает верхний элемент списка
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T // возвращаем нулевое значение для типа T
		return zero, false
	}

	last := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return last, true
}

func main() {
	stackInt := Stack[int]{}
	stackInt.Push(11)
	stackInt.Push(22)
	stackInt.Push(33)
	stackInt.Push(44)
	stackInt.Push(55)
	fmt.Println("stack after pushing several values:")
	fmt.Println(stackInt)

	fmt.Println(stackInt.Pop())
	fmt.Println(stackInt)
	fmt.Println(stackInt.Pop())
	fmt.Println(stackInt)

	stackStr := Stack[string]{}
	stackStr.Push("Huggy")
	stackStr.Push("Wuggy")
	stackStr.Push("Kissy")
	stackStr.Push("Missy")

	fmt.Println(stackStr)
	stackStr.Pop()
	fmt.Println(stackStr)
	stackStr.Pop()
	fmt.Println(stackStr)
}
