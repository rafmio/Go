package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	Height  int
}

func createStruct(n, s string, h int) *myStructure {
	if h > 300 {
		h = 0
	}

	return &myStructure{Name: n, Surname: s, Height: h}
}

func retStructure(n, s string, h int) myStructure {
	if h > 300 {
		h = 0
	}

	return myStructure{Name: n, Surname: s, Height: h}
}

func main() {
	s1 := createStruct("Mihalis", "Tsoukalos", 123)
	s2 := retStructure("Mihalis", "Tsoukalos", 123)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
}
