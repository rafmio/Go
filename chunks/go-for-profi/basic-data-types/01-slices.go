package main

import (
	"math/rand"
)

func fillSls(sls []int) {
	for i := 0; i < len(sls); i++ {
		randNum := rand.Intn(i+100)
		sls[i] = randNum
	}

	return
}

func changeSls(sls []int) {
	for _, value := range sls {
		value *= 10
	}

	return
}

func printSls(sls []int) {
	for i, val := range sls {
		println(i, ":", val)
	}
}

func main() {
	sls := make([]int, 10)
	fillSls(sls)
	changeSls(sls)
	printSls(sls)
}
