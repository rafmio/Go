// Determining whether a file exists
package main

import (
	"os"
)

func main() {
	targetFiles := []string{"no_such_file.txt", "MyTempFile.json", "printer.go"}
	for _, name := range targetFiles {
		info, err := os.Stat(name)
		if os.IsNotExist(err) {
			Printfln("File does not exits: %v, err: %v", name, err.Error())
		} else if err != nil {
			Printfln("Other error: %v", err.Error())
		} else {
			Printfln("File %v, Size: %v", info.Name(), info.Size())
		}

	}
}
