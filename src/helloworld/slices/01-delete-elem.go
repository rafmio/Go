package main

import "fmt"

func main() {
	idxForDel := 2 // "C"

	// Method 1: changes order (quick version)
	a := []string{"A", "B", "C", "D", "E"}

	a[idxForDel] = a[len(a)-1] // copy last element to the index to be deleted
	fmt.Println(a)
	a = a[:len(a)-1] // truncate the slice

	fmt.Println(a)

	// Method 2: saves order (low version)
	b := []string{"A", "B", "C", "D", "E"}
	copy(b[idxForDel:], b[idxForDel+1:]) // shift to the left by one idx
	b[len(b)-1] = ""                     // delete last element
	b = b[:len(b)-1]                     // truncate the slice

	fmt.Println(b)
}
