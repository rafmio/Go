// Locating files using a pattern
package main

import (
	"os"
	"path/filepath"
)

func main() {
	path, err := os.Getwd()
	Printfln("Getwd(): %v", path)
	Printfln("%s", "")

	if (err == nil) {
		matches, err := filepath.Glob(filepath.Join(path, "*.json"))
		Printfln("len(matches): %v", len(matches))
		if (err == nil) {
			for i, m := range matches {
				Printfln("%v elem - match: %v", i, m)
			}
		}
	}

	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}

	Printfln("%s", "")
	Printfln("%s", "pattern:")
	Printfln("%s", filepath.Join(path, "*.json"))
}

// func Glob(pattern string) (matches []string, err error)
