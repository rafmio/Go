// https://gobyexample.com/range-over-iterators
package main

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}
