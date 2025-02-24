package main

import (
	"fmt"
)

type tstStruct struct {
	mp  map[string]int
	num int
}

func main() {
	tst := new(tstStruct)
	fmt.Println(tst)

	fmt.Println(tst.mp)

	tst.mp = make(map[string]int)
	tst.mp["one"] = 1
	tst.mp["two"] = 2

	fmt.Println(tst.mp)

	fmt.Println(tst.num)
	tst.num = 10
	fmt.Println(tst.num)
}
