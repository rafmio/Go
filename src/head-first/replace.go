package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)

	fmt.Printf("type of broken: %T, %v\n", broken, broken)
	fmt.Printf("type of replace: %T, %v\n", replacer, replacer)
	fmt.Println(fixed)
}

// Functions are belong to packages, but methods are
// belong to concrete value
