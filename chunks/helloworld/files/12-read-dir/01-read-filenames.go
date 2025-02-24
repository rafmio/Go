package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("/home/raf/Go/src/helloworld/files/*")
	for _, f := range files {
		fmt.Println(f)
	}
}
