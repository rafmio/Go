package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func openFileFunc(file *os.File) error {
	var err error
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("scanning lines:", err.Error())
		// TODO: log error
	}

	filePosition, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		fmt.Println("setting file position:", err.Error())
		// TODO: log error
	}
	fmt.Println("file position inside openFileFunc:", filePosition)

	return err
}
