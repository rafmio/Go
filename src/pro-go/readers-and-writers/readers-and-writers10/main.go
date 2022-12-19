package main

import (
	// "bufio"
	// "io"
	"strings"
)

func main() {
	text := "It was a boat. A small boat"

	var builder strings.Builder
	var writer = NewCustomWriter(&builder)
	for i := 0; true; {
		end := i + 5
		if (end >= len(text)) {
			writer.Write([]byte(text[i:]))
			break
		}
		writer.Write([]byte(text[i : end]))
		i = end
	}

	Printfln("Writted data: %v", builder.String())
}
