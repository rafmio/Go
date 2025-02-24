package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example 1 -------------------------------------
	str1 := "!!Welcome to Golang tutorial !!"
	str2 := "@@This is the tutorial of Golang $$"

	fmt.Println("Strings before trimming: ")
	fmt.Println("str1: ", str1)
	fmt.Println("str2: ", str2)

	res1 := strings.Trim(str1, "!")
	res2 := strings.Trim(str2, "@$")
	fmt.Println()

	fmt.Println("Strings after trimming: ")
	fmt.Println("str1: ", res1)
	fmt.Println("str2: ", res2)

	fmt.Println("-------------------------------------")
	// Example 2 -------------------------------------
	str3 := "!!!Welcome to the world of Masterboy **"
	str4 := "@@ This is the tutorial of Golang$$"

	fmt.Println("String before trimming: ")
	fmt.Println("str3: ", str3)
	fmt.Println("str4: ", str4)
	fmt.Println()

	res3 := strings.TrimLeft(str3, "!*")
	res4 := strings.TrimLeft(str4, "@ ")

	fmt.Println("String after trimming: ")
	fmt.Println("res3: ", res3)
	fmt.Println("res4: ", res4)

	fmt.Println("-------------------------------------")
	// Example 3 -------------------------------------
	fmt.Println("Strings 3-4 before TrimRight: ")
	fmt.Println("str3: ", str3)
	fmt.Println("str4: ", str4)
	fmt.Println()

	res5 := strings.TrimRight(str3, "!*")
	res6 := strings.TrimRight(str4, "$")

	fmt.Println("Strings 3-4 after TrimRight: ")
	fmt.Println("res5: ", res5)
	fmt.Println("res6: ", res6)

	fmt.Println("-------------------------------------")
	// Example 4 -------------------------------------
	fmt.Println("Strings 3-4 before TrimSpace: ")
	fmt.Println("str3: ", str3)
	fmt.Println("str4: ", str4)
	fmt.Println()

	res7 := strings.TrimSpace(str3)
	res8 := strings.TrimSpace(str4)

	fmt.Println("String 3-4 after TrimSpace: ")
	fmt.Println("str3: ", res7)
	fmt.Println("str4: ", res8)

	fmt.Println("-------------------------------------")
	// Example 5 -------------------------------------
	str5 := "First thing first"
	str6 := "Take a look around"

	fmt.Println("Strings before trimming: ")
	fmt.Println("str5: ", str5)
	fmt.Println("str6: ", str6)
	fmt.Println()

	res9 := strings.TrimSuffix(str5, "first")
	res10 := strings.TrimSuffix(str6, "look")

	fmt.Println("Strings after trimming: ")
	fmt.Println("res9:  ", res9)
	fmt.Println("res10: ", res10)

	fmt.Println("-------------------------------------")
	// Example 6 -------------------------------------
	str11 := "First thing first"
	str12 := "Take a look around"

	fmt.Println("Strings before trimming: ")
	fmt.Println("str11: ", str11)
	fmt.Println("str12: ", str12)
	fmt.Println()

	res11 := strings.TrimPrefix(str11, "First")
	res12 := strings.TrimPrefix(str12, "Hello")

	fmt.Println("Strings after trimming: ")
	fmt.Println("res11: ", res11)
	fmt.Println("res12: ", res12)

}

// Trim a String
// https://www.geeksforgeeks.org/how-to-trim-a-string-in-golang/
