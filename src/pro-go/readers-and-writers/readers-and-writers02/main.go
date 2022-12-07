package main

import (
	"io"
	"strings"
)
func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if (count > 0) {
			writer.Write(b[0 : count])
			Printfln("Read %v bytes: %v", count, string(b[0 : count]))
			// Printfln("Type of writer: %T", writer)
		}
		if (err == io.EOF) {
			break
		}
	}
}

func main() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	// Printfln("Type of builder: %T", builder)

	processData(r, &builder)

	Printfln("String builder contents: %s", builder.String())
}


// https://cs.opensource.google/go/go/+/refs/tags/go1.19.4:src/strings/reader.go;l=17;bpv=1;bpt=1
// string.Reader struct:
//
// type Reader struct {
// 	s        string
// 	i        int64 // current reading index
// 	prevRune int   // index of previous rune; or < 0
// }


// https://cs.opensource.google/go/go/+/refs/tags/go1.19.4:src/strings/builder.go;l=15
// string.Builder struct:
//
// type Builder struct {
// 	addr *Builder // of receiver, to detect copies by value
// 	buf  []byte
// }
