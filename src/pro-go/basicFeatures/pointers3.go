package main

import (
	"fmt"
	"sort"
)

func main() {
	names := [3]string{"Alice", "Charlie", "Bob"}
	secondName := names[1]
	fmt.Printf("Type of secondName: %T\n", secondName)
	fmt.Println(secondName)
	sort.Strings(names[:])
	fmt.Println(secondName)
	fmt.Println()

	secondPosition := &names[1]
	fmt.Printf("Type of secondPosition: %T\n", secondPosition)
	fmt.Println(*secondPosition)
	sort.Strings(names[:])
	fmt.Println(*secondPosition)

}
