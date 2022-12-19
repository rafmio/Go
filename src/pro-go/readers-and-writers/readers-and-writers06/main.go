package main

import (
	"io"
	"strings"
)

func main() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Ball")
	r3 := strings.NewReader("Backet")

	concatReader := io.MultiReader(r1, r2, r3)

	var writer strings.Builder
	teeReader := io.TeeReader(concatReader, &writer)

	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())

	Printfln("%s", "")

	r4 := strings.NewReader("Kissy")
	r5 := strings.NewReader("Missy")
	r6 := strings.NewReader("Tosy")

	concatReader2 := io.MultiReader(r4, r5, r6)

	limited := io.LimitReader(concatReader2, 5)
	ConsumeData(limited)
}
