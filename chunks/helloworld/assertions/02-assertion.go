package main

import (
	"fmt"
)

type MyType struct {
	Value string
}

func main() {
	var i interface{} = "hello-mello"

	s := i.(string)
	fmt.Printf("Type: %T, Value: %s\n", s, s)

	// type assertion with checking
	s, ok := i.(string)
	if ok {
		fmt.Println("Assertion successful:", s)
		fmt.Printf("Type of s: %T, value of s: %v\n", s, s)
	} else {
		fmt.Println("Assertion failed")
	}

	// type assertion with wrong type
	var ii interface{} = 123
	s2, ok := ii.(string)
	if ok {
		fmt.Println("Assertion successful:", s2)
	} else {
		fmt.Println("Assertion failed")
		fmt.Printf("Type of ii: %T, value of ii: %v\n", ii, ii)
	}

	var iii interface{} = []byte("Hello-mello, tosy-bosy")

	if s3, ok := iii.([]byte); ok {
		fmt.Println("Assertion successful:", string(s3))
		fmt.Printf("Type of s3: %T, value of s3: %v\n", s3, s3)
	} else {
		fmt.Println("Assertion failed")
		fmt.Printf("Type of iii: %T, value of iii: %v\n", iii, iii)
	}

	var iiii interface{} = MyType{Value: "Huggy-Wuggy"}

	if s4, ok := iiii.(MyType); ok {
		fmt.Println("Assertion successful:", s4.Value)
		fmt.Printf("Type of s4: %T, value of s4: %v\n", s4, s4)
	} else {
		fmt.Println("Assertion failed")
	}
}
