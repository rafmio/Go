package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example 1 ------------------------------------
	str1 := "Welcome, to the, online portal, of GFG"
	str2 := "My dog's name is Luna"
	str3 := "I like to play guitar"

	fmt.Println("str1: ", str1)
	fmt.Println("str2: ", str2)
	fmt.Println("str3: ", str3)
	fmt.Println()

	res1 := strings.Split(str1, ",")
	res2 := strings.Split(str2, "")
	res3 := strings.Split(str3, "!")
	res4 := strings.Split("", "GeekForGeeks, geeks")

	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
	fmt.Println("res3: ", res3)
	fmt.Println("res4: ", res4)
	fmt.Println()

	for i := 0; i < len(res1); i++ {
		fmt.Printf("Index: %d, value: %s\n", i, res1[i])
	}

	fmt.Println("--------------------------------------------")
	// Example 2 ---------------------------------------------
	str4 := "Huggy-Wuggy, Kissy-Missy, Toddy-Boddy"
	str5 := "There is no reason to don't like Elvis Presley"
	str6 := "I like to play guitar"

	fmt.Println("str4: ", str4)
	fmt.Println("str5: ", str5)
	fmt.Println("str6: ", str6)
	fmt.Println()

	res5 := strings.SplitAfter(str4, ",")
	res6 := strings.SplitAfter(str5, "")
	res7 := strings.SplitAfter(str6, "!")
	res8 := strings.SplitAfter("", "GeeksforGeeks, geeks")

	fmt.Println("res5: ", res5)
	fmt.Println("res6: ", res6)
	fmt.Println("res7: ", res7)
	fmt.Println("res8: ", res8)

	for i := 0; i < len(res7); i++ {
		fmt.Printf("Index: %d, value: %s\n", i, res7[i])
	}

	fmt.Println("--------------------------------------------")
	// Example 3 ---------------------------------------------
	str8 := "Kissy-Missy, Huggy-Wuggy, Toddy-Buddy"
	str9 := "There is to reason to don't like Pink Floyd"
	str10 := "I like to play the guitar"

	fmt.Println("Before strings.SplitAfterN: ")
	fmt.Println("str8:  ", str8)
	fmt.Println("str9:  ", str9)
	fmt.Println("str10: ", str10)

	res9 := strings.SplitAfterN(str8, ",", 2)
	res10 := strings.SplitAfterN(str9, "", 4)
	res11 := strings.SplitAfterN(str10, "!", 1)
	res12 := strings.SplitAfterN("", "LimpBizkit, band", 3)

	fmt.Println("After strings.SplitAgterN: ")
	fmt.Println("res9:  ", res9)
	fmt.Println("res10: ", res10)
	fmt.Println("res11: ", res11)
	fmt.Println("res12: ", res12)
}
