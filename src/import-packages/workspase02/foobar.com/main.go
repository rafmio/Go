// https://linguinecode.com/post/how-to-import-local-files-packages-in-golang#step1
package main

import (
	// "github.com/rleija/go-test-project/woosah"
	"fmt"

	"go-test-project/woosah"
)

func main() {
	// cannot refer to  unexported name woosah.test1
	// fmt.Println(woosah.test1)

	// public variable
	fmt.Println(woosah.Test2)
}
