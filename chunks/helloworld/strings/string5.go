package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example 1 ----------------------------------
	str1 := "Grisley"
	str2 := "White-Bear"
	str3 := "Winnie-Pooh"
	str4 := "Grisley"

	fmt.Println("str1 == str2: ", str1 == str2)
	fmt.Println("str2 == str3: ", str2 == str3)
	fmt.Println("str3 == str4: ", str3 == str4)
	fmt.Println("str4 == str1: ", str4 == str1)
	fmt.Println()

	fmt.Println("str1 != str2: ", str1 != str2)
	fmt.Println("str2 != str3: ", str2 != str3)
	fmt.Println("str3 != str4: ", str3 != str4)
	fmt.Println("str4 != str1: ", str4 != str1)

	fmt.Println("----------------------------------")
	// Example 2 ----------------------------------
	fmt.Println(strings.Compare("Bear", "White Bear"))
	fmt.Println(strings.Compare("Grisley", "Grisley"))
	fmt.Println(strings.Compare("White Bear", "Bear"))
}

// Ways to compare strings
// https://www.geeksforgeeks.org/different-ways-to-compare-strings-in-golang/
