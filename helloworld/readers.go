package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader")
	fmt.Println(r)

	b := make([]byte, 8)

	fmt.Printf("Type of r: %T, type of b: %T", r, b)
	fmt.Println()

	fmt.Println("r: ", r)
	fmt.Println("b: ", b)

	fmt.Println()

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}

// Readers
// https://go.dev/tour/methods/21
