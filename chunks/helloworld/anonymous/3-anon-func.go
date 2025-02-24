// https://kovardin.ru/articles/go/zamykaniya/
package main

import (
	"fmt"
)

func outer(name string) {
	text := "Modified " + name
	foo := func() {
		fmt.Println(text)
	}

	foo()
}

func main() {
	outer("hello-mello")
}
