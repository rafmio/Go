package main

import (
	"fmt"
	"log"
)

func main() {
	a := 100
	b := 200
	var err error

	if a < b {
		fmt.Println("a is less than b")
		err = fmt.Errorf("a is less than b")
	}

	log.Printf("%d is greater than %d, and error is %v", b, a, err)
}
