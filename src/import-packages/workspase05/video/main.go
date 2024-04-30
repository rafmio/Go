package main

import (
	"fmt"
	"video/testPackage"
	"video/testPackage2"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(testPackage.Greet())
	testPackage2.Mew()
}
