package main

import "fmt"

func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])

	var primes [5]int = [5]int{2, 4, 6, 10, 127}
	fmt.Println(primes[0], primes[2], primes[4])

	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}

	notes2 := [6]string{"mi", "si", "so", "re", "la", "mi"}
	fmt.Println(notes2[2], notes2[4])

	fmt.Println()

	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line",
	}

	fmt.Println(text)
	fmt.Println(text[2])
}
