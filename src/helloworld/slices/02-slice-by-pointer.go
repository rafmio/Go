package main

import (
	"fmt"
	"math/rand"
)

func changeSls(sls []int) {
	sls[0] = 100
	sls = append(sls, 1000) // the slice will not be changed because append returns a new slice
}

func main() {
	sls := make([]int, 5)

	for i := 0; i < len(sls); i++ {
		sls[i] = rand.Intn(100)
	}

	fmt.Println(sls)

	changeSls(sls)

	fmt.Println(sls)
}
