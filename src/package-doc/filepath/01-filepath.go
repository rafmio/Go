package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dirPath := "/app/input-data"
	fileName := "02-events.json"

	filePath := filepath.Join(dirPath, fileName)

	fmt.Println("File path:", filePath)
}
