package main

import (
	"fmt"
	"math/rand"
)

func fillSls(sls []int) {
	for i := 0; i < len(sls); i++ {
		randNum := rand.Intn(i+100)
		sls[i] = randNum
	}
}

func printMOS(mos map[string]string) {
	for key, val := range mos {
		fmt.Println(key, ":", val)
	}
}

func main() {
	mapOfSls := make(map[string]string)
	s1 := make([]int, 10)
	fillSls(s1)

	mapOfSls["s1"] = fmt.Sprintf("%p", s1)

	s2 := s1[3:7]
	mapOfSls["s2"] = fmt.Sprintf("%p", s2)

	s3 := s1[2:8]
	mapOfSls["s3"] = fmt.Sprintf("%p", s3)

	ss2 := s2[2:3]
	mapOfSls["ss2"] = fmt.Sprintf("%p", ss2)

	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	printMOS(mapOfSls)

	fmt.Println("s1:", s1)
	fmt.Printf("s1's addr: %p\n", s1)
}
