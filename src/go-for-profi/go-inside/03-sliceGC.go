package main

import "runtime"

type data struct {
	i, j int
}

func main() {
	var N = 40_000_000
	var slsOfData []data

	for i := 0; i < N; i++ {
		value := int(i)
		slsOfData = append(slsOfData, data{i, value + i})
	}

	runtime.GC()
	_ = slsOfData[0]
}
