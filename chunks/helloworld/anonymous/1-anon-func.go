package main

import (
	"fmt"
)

func GFG() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "Paul McCartney"
	}

	return myf
}

func main() {
	value := GFG()
	fmt.Println(value("Welcome ", "to "))

	myf2 := func(i int) string {
		var str string
		str = fmt.Sprintf("Now val is: %d", i)

		return str
	}

	for i := 0; i < 5; i++ {
		fmt.Println(myf2(i))
	}
}
