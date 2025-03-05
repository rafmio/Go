package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Printf("Hello-mello, %s, this function will be wrapped\n", name)
}

func wrappedGreet(name string) {
	fmt.Printf("Before greeting\n")
	greet(name)
	fmt.Printf("After greeting\n")
}

func main() {
	wrappedGreet("John Doe")
}
