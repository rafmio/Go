package main

import (
	"fmt"
	"math/rand"
)

// Stack - stack for any data type
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, true
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func main() {
	intStack := Stack[int]{}
	fmt.Printf("Type of intStack: %T\n", intStack)

	stcSize := rand.Intn(10)
	for range stcSize {
		intStack.Push(rand.Intn(100))
	}

	fmt.Println("Int stack size:", intStack.Size())

	for intStack.Size() > 0 {
		item, _ := intStack.Pop()
		fmt.Printf("%d - %T\n", item, item)
	}

	fmt.Println()

	stringStack := Stack[string]{}
	stringStack.Push("Go")
	stringStack.Push("is")
	stringStack.Push("awesome!")
	stringStack.Push("Tosy")
	stringStack.Push("Bosy")
	fmt.Printf("Type of stringStack: %T\n", stringStack)
	fmt.Println("string stack size:", stringStack.Size())

	for stringStack.Size() > 0 {
		item, _ := stringStack.Pop()
		fmt.Printf("%s - %T\n", item, item)
	}
}
