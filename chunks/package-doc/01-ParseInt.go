package main

import (
	"fmt"
	"strconv"
)

func main() {
	runParseInt()
	runParseHex()
	runAtoi()
}

func runParseInt() {
	str := "12345"
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("Parsed integer: %d, type: %T\n", i, i)
	}
}

func runParseHex() {
	str := "077"
	i, err := strconv.ParseInt(str, 8, 64) // should return '63'
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("Parsed integer: %d, type: %T\n", i, i)
	}
}

func runAtoi() {
	str := "42"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("Parsed integer: %d, type: %T\n", i, i)
	}
}
