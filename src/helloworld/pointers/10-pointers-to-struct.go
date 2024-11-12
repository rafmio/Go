package main

import (
	"fmt"
)

type MyStruct struct {
	Name   string
	Age    uint
	Height uint
}

func main() {
	ptrInst := new(MyStruct)
	ptrInst.Name = "Andy"
	ptrInst.Age = 41
	ptrInst.Height = 86

	fmt.Println("ptrInst:", ptrInst)

	inst := MyStruct {
		Name: "Jerry",
		Age: 26,
		Height: 76,
	}

	fmt.Println("ints:", inst)

	ptrInst = &inst

	fmt.Println("ptrInst:", ptrInst)

	return
}
