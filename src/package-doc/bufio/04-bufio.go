// https://metanit.com/go/tutorial/8.9.php
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rows := []string{
		"hello-mello",
		"Huggy-Wuggy",
		"Kissy-Missy",
		"Tosy-Bosy",
	}

	file, err := os.Create("some.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer file.Close()

	// bufio.NewWriter returns *Writer
	// Writer - то, куда записываем
	writer := bufio.NewWriter(file)

	for _, row := range rows {
		writer.WriteString(row)
		writer.WriteString("\n")
	}

	// ? до вызова bufio.Flush() все записанные данные находятся в памяти
	// после вызова - сбрасываются на диск
	writer.Flush()
}

// Write to file
// 1. create or open file to var 'file' ('file' implements 'io.Writer' interface)
// 2. create new bufio.Writer via 'bufio.NewWriter' func, passing file as an argument
// 3. the 'bufio.NewWriter' func takes io.Writer ('file' var) as an argument
// 4. Write strings via 'bufio.WriteStrings' func using 'for' loop
