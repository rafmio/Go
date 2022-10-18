package main

import "fmt"

func main() {
	var notes [3]string = [3]string{"do", "re", "mi"}
	var primes [3]int = [3]int{2, 3, 122}

	fmt.Println(notes)
	fmt.Println(primes)

	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)
}
