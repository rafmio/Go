// https://habr.com/ru/post/490336/

package main

import (
	"fmt"
	"os"
)

func printIntVal(c chan int) {
	var outvalue int
	outvalue = <-c
	fmt.Println("outvalue: ", outvalue)
}

func main() {
	c := make(chan int)

	fmt.Printf("Type of `c` is: %T\n", c)
	fmt.Printf("Value of `c` is: %v\n", c)

	var intvalue int = 100
	go printIntVal(c)
	c <- intvalue
	close(c)

	os.Exit(0)
}
