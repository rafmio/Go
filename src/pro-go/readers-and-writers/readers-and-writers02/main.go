package main

import (
	"io"
	"strings"
	// "fmt"
)

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}

		if err == io.EOF {
			break
		}
	}
}

func main() {
	r := strings.NewReader("Kayak")
	Printfln("Type of r is: %T", r)
	var builder strings.Builder
	Printfln("Type of builder is: %T", builder)
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())

}

// strings.Builder - is used to efficently build a string using Write methods
