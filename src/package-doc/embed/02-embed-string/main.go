// https://www.sobyte.net/post/2021-06/go-embed-brief-tutorial/
package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var s string

func main() {
	fmt.Println(s)
}
