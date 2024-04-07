// https://www.sobyte.net/post/2021-06/go-embed-brief-tutorial/
package main

import (
	"embed"
	"fmt"
)

//go:embed p/hello2.txt
//go:embed p/hello1.txt
var f embed.FS

func main() {
	data, _ := f.ReadFile("p/hello1.txt")
	fmt.Println(string(data))

	data2, _ := f.ReadFile("p/hello2.txt")
	fmt.Println(string(data2))
}
