package main

import (
	"fmt"
)

type MyMap map[string]int

func (m MyMap) Length() int {
	return len(m)
}

func (m MyMap) Sum() int {
	var sum int
	for _, val := range m {
		sum += val
	}
	return sum
}

func main() {
	mm := MyMap{"A": 1, "B": 2}

	fmt.Println(mm)
	fmt.Println(mm["A"])
	fmt.Println(mm.Length())

	mmm := MyMap{"A": 1, "B": 2, "C": 3}
	fmt.Println(mm.Sum())
	fmt.Println(mmm.Sum())
}
