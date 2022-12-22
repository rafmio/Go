// Enumerating files
package main

import (
	"os"
)

func main() {
	path, err := os.Getwd()
	Printfln("path: %v", path)
	Printfln("%s", "")
	if (err == nil) {
		dirEntries, err := os.ReadDir(path)
		if (err == nil) {
			for _, dentry := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
			}
		}
	}

	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}
