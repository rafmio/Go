package main

import (
	"io"
	"os"
	"log"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(os.Args) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "This is Standard output\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stdout, "\n")

	log.Println("This is log.Println() string")
}
