package main

import (
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	println("*p1: ", *p1)
	println("*p2: ", *p2)
	*p1 = 5434123412312431212
	println("value:", value)
	println("*p2: ", *p2)
	*p1 = 54341234
	println("value:", value)
	println("*p2: ", *p2)
}
