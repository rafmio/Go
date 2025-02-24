package main

import (
	"fmt"
	"unicode/utf8" // For Example 4
)

func main() {
	// Example 1 ----------------------------------------
	my_value1 := "Welcome to world of Scatman John"

	var my_value2 string
	my_value2 = "Scatman John"

	fmt.Println("my_value1: ", my_value1)
	fmt.Println("my_value2: ", my_value2)

	fmt.Println("---------------------------------------")
	// Example 2 ----------------------------------------
	my_value_3 := "Come down throught..."
	my_value_4 := "Come \ndown \nthrought..."
	my_value_5 := `Create string using backticks`
	my_value_6 := `String created \nusing backticks`

	fmt.Println("my_value_3: ", my_value_3)
	fmt.Println("my_value_4: ", my_value_4)
	fmt.Println("my_value_5: ", my_value_5)
	fmt.Println("my_value_6: ", my_value_6)

	fmt.Println("---------------------------------------")
	// Example 2 ----------------------------------------
	// Iterating strings
	for index, s := range "Scatman John" {
		fmt.Printf("The index number of %c is %d\n", s, index)
	}

	fmt.Println("---------------------------------------")
	// Example 3 ----------------------------------------
	// Create a string from the slice
	slc1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	my_string1 := string(slc1)

	fmt.Println("my_string1: ", my_string1)

	slc2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	my_string2 := string(slc2)

	fmt.Println("my_string2: ", my_string2)

	fmt.Println("---------------------------------------")
	// Example 4 ----------------------------------------
	// Len of the string

	my_string3 := "Huggy Wuggy Kissy Missy"
	length1 := len(my_string3)
	length2 := utf8.RuneCountInString(my_string3)

	fmt.Println("String: ", my_string3)
	fmt.Println("Length 1: ", length1)
	fmt.Println("Length 2: ", length2)
}
