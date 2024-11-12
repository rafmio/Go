package main

import (
	"fmt"
)

type MyStruct struct {
	Name string
	Age uint
	Salary uint
	Height uint
}

func main() {
	Ni := newInstance()	
	fmt.Printf("Type of ni: %T\n", Ni)

	fmt.Println(Ni)

	Ni.Salary = 1_540_000
	fmt.Println(Ni)

	Ni.Height = 180
	fmt.Println(Ni)
}

func newInstance() *MyStruct {
	ni := new(MyStruct)
	ni.Name = "Larry"
	ni.Age = 41

	return ni
}
