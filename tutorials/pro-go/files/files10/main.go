// Managing files and directories
package main

import (
	"os"
	"encoding/json"
	"path/filepath"
	"fmt"
)

func main() {
	path, err := os.UserHomeDir()
	if (err == nil) {
		path = filepath.Join(path, "go/src/pro-go/files/files10", "MyTempFile.json")
	}

	Printfln("Full path: %v", path)

	err = os.MkdirAll(filepath.Dir(path), 0766)
	if (err == nil) {
		file, err := os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0666)
		if (err == nil) {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
		if (err != nil) {
			Printfln("Error: %v", err.Error())
		}
	}

	err = os.MkdirAll("./newdir", 0766)
	if (err != nil) {
		fmt.Println(err.Error)
	} else {
		fmt.Println("New dir was created")
	}
}
