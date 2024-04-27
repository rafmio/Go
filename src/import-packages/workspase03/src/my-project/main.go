package main

import (
	"fmt"
	foo "my_project/foo"
)

func main() {
	fmt.Println(foo.TstString)
	foo.TryToImport()
}
