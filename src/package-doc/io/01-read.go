// https://golangdocs.com/io-package-in-golang
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Read this string")
	fmt.Printf("Type of r: %T\n", r)

	buf := make([]byte, 20)
	numRead, err := io.ReadAtLeast(r, buf, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("numRead:", numRead)
	}

	fmt.Println(string(buf))

	buf = make([]byte, 2)
	numRead, err = io.ReadAtLeast(r, buf, 4)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("numRead:", numRead)
	}
	fmt.Println(string(buf))

	buf = make([]byte, 20)
	numRead, err = io.ReadAtLeast(r, buf, 20)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Println("numRead:", numRead)
	}
	fmt.Println(string(buf))
}
