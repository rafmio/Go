// https://www.sobyte.net/post/2021-06/go-embed-brief-tutorial/
package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed hello.txt
//go:embed hello2.txt
var f embed.FS

func main() {
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))

	data2, _ := f.ReadFile("hello2.txt")
	fmt.Println(string(data2))
}
