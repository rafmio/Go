package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Enter arguments")
	}
	arglist := os.Args

	var sumOfArgs float64
	sumOfArgs = 0

	for _, argument := range arglist {
		value, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			fmt.Println("parsing: ", err.Error())
		} else {
			sumOfArgs += value
		}
	}

	fmt.Println("Sum of all arguments: ", sumOfArgs)
	fmt.Printf("Average of all float values: %.2f\n", sumOfArgs/float64(len(arglist)-1))
}
