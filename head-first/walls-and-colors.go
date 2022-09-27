package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("a width %0.2f or height %0.2f are invalid", width, height)
	} else {
		area := width * height
		return area / 10.0, nil
	}
}

func main() {
	amount, err := paintNeeded(4.2, 3.0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%0.2f liters needed\n", amount)

	fmt.Println()
}
