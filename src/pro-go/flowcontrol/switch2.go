package main

import "fmt"

func main() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", counter/2)
		default:
			fmt.Println("Non-prime value:", counter/2)
		}
	}
	fmt.Println()
	// The duplication can be avoided using an initialization statement:
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value:", val)
		default:
			fmt.Println("Non-prime value:", val)
		}
	}
	fmt.Println()
	// Omitting a Comparison Value
	// When the comparison value is omitted, each case statment is specified
	// with a condition
	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}
}
