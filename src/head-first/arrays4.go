package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println("index: ", index, notes[index])
	index = 3
	fmt.Println("index: ", index, notes[index])

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	fmt.Println()
	fmt.Println("for..range:")
	for index, note := range notes {
		fmt.Println(index, note)
	}

	fmt.Println()
	fmt.Println("using _ underscore")

	for _, note := range notes {
		fmt.Println(note)
	}
}
