package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		os.Exit(1)
	}

	// открываем файл для чтения/записи файловой позиции
	filePoisitionFile, err := os.Open("filePosition")
	if err != nil {
		if err == os.ErrNotExist {
			filePoisitionFile, err = os.Create("filePosition")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// открываем файл для чтения записей логов
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// находим файловую позицию
	filePoisition, err := file.Seek(0, os.SEEK_CUR)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
