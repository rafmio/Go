package main

import (
	"io"
	"os"
)

func main() {
	myStrinig := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myStrinig = "Please give me one argument!"
	} else {
		myStrinig = "Hello, " + arguments[1] + "!"
	}

	io.WriteString(os.Stdout, myStrinig)
	io.WriteString(os.Stdout, "\n")
}
