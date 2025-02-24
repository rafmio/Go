package main

import (
	"fmt"
)

func main() {
	var names2 = make(map[int]string)
	names2[0] = "John"
	names2[1] = "Jane"
	names2[2] = "Jerry"
	fmt.Println(names2)

	var names3 = map[int]string{
		0: "Jessy",
		1: "Walter",
		2: "Hank",
	}
	fmt.Println(names3)

	var names4 = make(map[int]string)
	names4[0] = "Tukko"
	names4[1] = "Elladio"
	names4[2] = "Hector"
	names4[3] = "Soul"
	fmt.Println(names4)

	for i, val := range names4 {
		fmt.Printf("%d: %s\n", i, val)
	}

	checkName, exists := names4[0]
	if exists {
		fmt.Printf("%s exists\n", checkName)
	} else {
		fmt.Println("Doesn't exist")
	}

}
