package main

import "fmt"

type Slice []string

func Combination(s Slice, k int) <-chan Slice {
	channel := make(chan Slice)
	go func() {
		defer close(channel)

		if k <= 0 {
			channel <- Slice{}
			return
		}

		sls := Slice{}
		for i := 0; i < len(s); i++ {
			sls = Slice{s[i]}
			for anothSls := range Combination(s[i+1:], k-1) {
				channel <- append(sls, anothSls...)
			}
		}
	}()

	return channel
}

func main() {
	s := Slice{"a", "b", "c"}
	for sls := range Combination(s, 2) {
		fmt.Println(sls)
	}
}
