package main

import "fmt"

func main() {
	keys := []string{}
	carta := map[string][]int{
		"A": {1, 2},
		"B": {2, 10},
	}

	for key, _ := range carta {
		fmt.Printf("Current key: %s, type of key: %T\n", key, key)
		keys = append(keys, key)
	}

	fmt.Println(keys)

	for key, value := range carta {
		fmt.Println(key, value)
	}
}
