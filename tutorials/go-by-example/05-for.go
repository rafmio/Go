package main

import "fmt"

func main() {
	i := 1

	fmt.Println("------")
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("------")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	fmt.Println("range:")
	for i := range 3 {
		fmt.Println("range", i)
	}

	fmt.Println("------")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("------")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
