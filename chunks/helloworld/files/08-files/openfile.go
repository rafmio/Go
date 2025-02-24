package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

func main() {
	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Println("opening file:", err.Error())
	// 	// TODO: log error
	// }
	// defer file.Close()

	file, err := os.OpenFile("input.txt", os.O_RDONLY, fs.FileMode(os.O_EXCL))
	if err != nil {
		fmt.Println("opening file:", err.Error())
		// TODO: log error
	}
	defer file.Close()

	filePosition, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		fmt.Println("setting file position:", err.Error())
		// TODO: log error
	}
	fmt.Println("file position before:", filePosition)

	err = openFileFunc(file)
	if err != nil {
		fmt.Println(err)
	}

	filePosition, err = file.Seek(0, io.SeekCurrent)
	if err != nil {
		fmt.Println("setting file position:", err.Error())
		// TODO: log error
	}
	fmt.Println("file position after:", filePosition)

	time.Sleep(time.Second * 10)

	err = openFileFunc(file)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * 10)
}
