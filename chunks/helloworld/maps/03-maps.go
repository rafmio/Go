package main

import "fmt"

func main() {
	products := map[string]float64{
		"Java":       279,
		"JavaScript": 49.95,
	}

	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Java"])
	fmt.Println("Price:", products["JavaScript"])
}
