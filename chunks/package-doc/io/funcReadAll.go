// DOES NOT WORK! (io.ReadAll undefined) go version 1.15.15
// https://pkg.go.dev/io
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}
// func ReadAll(r Reader) ([]byte, error)
// ReadAll reads r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF
// ReadAll is defined to read from src until EOF, it does not treat EOF as an error
