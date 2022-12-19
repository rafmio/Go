package main

import (
	"io"
	"strings"
)

func main() {
	pipeReader, pipeWriter := io.Pipe()
	// go func() {
	//   GenerateData(pipeWriter)
	//   pipeWriter.Close()
	// }()

	go GenerateData(pipeWriter)

	ConsumeData(pipeReader)
	Printfln("%s", "")
	// ---------------------------------------------------------
	// Concatenating Multiple Readers
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	r4 := strings.NewReader("Nuckear Submarine")

	concatReader := io.MultiReader(r1, r2, r3, r4)

	ConsumeData(concatReader)

	Printfln("%s", "")
	// ---------------------------------------------------------
	// Combining Multiple Writers
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder

	combinedWriter := io.MultiWriter(&w1, &w2, &w3)

	GenerateData(combinedWriter)

	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())
}
