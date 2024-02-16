package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var s string = "Hello-mello, string-mring"
	var b []byte = []byte{90, 104, 106, 20, 107, 108, 109, 20, 110, 111, 112}

	readerS := bytes.NewReader([]byte(s))
	readerB := bytes.NewReader(b)

	fmt.Printf("Type of readerS: %T\n", readerS)
	fmt.Printf("Type of readerB: %T\n", readerB)

	output(readerS)
	fmt.Println()
	output(readerB)
}

func output(reader *bytes.Reader) {
	buf := make([]byte, 40)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading:", err.Error())
		} else if n > 0 {
			fmt.Println(buf[:n])
			fmt.Println(string(buf[:n]))
		}
	}
}
