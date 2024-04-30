package main

import (
	"example/user/hello/morestrings"
	"fmt"
)

func main() {
	fmt.Println("Hello, world! Hey!")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(morestrings.ReverseRunes("ygguW ygguH"))
}
