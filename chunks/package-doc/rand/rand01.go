package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var charset = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+"
	for _, value := range charset {
		fmt.Printf("%s-", string(value))
	}

	fmt.Println(charset)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println()

	var password [30]string
	// for i, value := range password {
	// 	randInx := rand.Intn(len(charset) - 1)
	// 	value = string(charset[randInx])
	// 	fmt.Printf("i: %d, %s\n", i, value)
	// }

	for i := 0; i < 30; i++ {
		randInx := rand.Intn(len(charset) - 1)
		password[i] = string(charset[randInx])
	}

	fmt.Println("!--------------------------------!")

	var passwordStr string
	for i := 0; i < 30; i++ {
		passwordStr += password[i]
	}

	fmt.Println(passwordStr)
}
