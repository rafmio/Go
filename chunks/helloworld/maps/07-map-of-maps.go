package main

import (
	"fmt"
)

func main() {
	mpOfMaps := make(map[string]map[string]int)

	mpOfMaps["a"] = make(map[string]int)
	mpOfMaps["b"] = make(map[string]int)

	mpOfMaps["a"]["a1"] = 1
	mpOfMaps["b"]["b1"] = 2

	fmt.Println(mpOfMaps)
}
