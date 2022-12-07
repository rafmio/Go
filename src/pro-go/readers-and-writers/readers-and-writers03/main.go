package main

import (
	"io"
	"strings"
)

func processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if(err == nil) {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	r := strings.NewReader("Kayak") // Returns pointer
	var builder strings.Builder // Builder is regular var, don't pointer
	Printfln("Type of builder: %T", builder)
	processData(r, &builder)
	Printfln("String builder contents: %s", builder.String())
}

// https://cs.opensource.google/go/go/+/refs/tags/go1.19.4:src/strings/builder.go;l=47
// String returns the accumulated string.
// func (b *Builder) String() string {
// 	return *(*string)(unsafe.Pointer(&b.buf))
// }
