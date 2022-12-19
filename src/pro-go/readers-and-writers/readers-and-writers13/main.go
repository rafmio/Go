// Writing formatted string to a writer
package main

import (
	"io"
	"strings"
	"fmt"
)

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func main() {
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"

	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))

	fmt.Println(writer.String())
	fmt.Println()
	// -------------------------------------------------------------------------

	var writer2 strings.Builder
	fmt.Fprintf(&writer2, template, "Submarine", "Warship games", float64(1_000_000))
	fmt.Println(writer2.String())
}
