package main

import "fmt"

func main() {
	fmt.Printf("Please, enter the number: ")
	var number int
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Printf("%d is a factor of %d\n", i, number)
		}
	}
}
