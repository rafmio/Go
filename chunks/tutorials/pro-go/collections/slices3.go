package main

import "fmt"

func main() {
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"

	moreNames := []string{"Hat Gloves"}
	appendNames := append(names, moreNames...)
	fmt.Println("appendNames", appendNames)
	fmt.Println()

	var products [4]string
	for index, _ := range products {
		products[index] = appendNames[index]
	}
	fmt.Println("products: ", products)
	someNames := products[1:3]
	allNames := products[:]

	fmt.Println("someNames:", someNames)
	fmt.Println("allNames:", allNames)
	for index, _ := range products {
		fmt.Println("products == allNames:", products[index] == allNames[index])
	}
}
