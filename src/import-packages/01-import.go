package main

import (
	"fmt"

	"./dir/otherpackage"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello,", Name)
	fmt.Println("Hello,", otherpackage.NameFormOtherPackage)
}
