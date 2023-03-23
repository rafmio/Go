package main

import (
	"fmt"
)

const (
	Red int = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

func main() {
	fmt.Printf("Red: %d\nOrange: %d\nYello: %d\nGreen: %d\nBlue: %d\nIndigo: %d\nViolet: %d\n",
		Red, Orange, Yellow, Green, Blue, Indigo, Violet,
	)
}
