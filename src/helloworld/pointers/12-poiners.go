package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	num := 100
	numPtr := &num

	fmt.Println(num)
	_ = passViaValue(num)
	fmt.Println(num)
	passViaPointer(numPtr)
	fmt.Println(num)
}

func passViaValue(num int) int {
	result := num * 20

	return result
}

func passViaPointer(numPtr *int) {
	*numPtr = *numPtr * 20

}
