package main

import (
	"bufio"
	"fmt"
	"go/scanner"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		os.Exit(1)
	}

	// открываем файл для чтения/записи файловой позиции
	filePositionFile, err := os.Open("filePosition")
	if err != nil {
		if err == os.ErrNotExist {
			// создаём файл, если он не существует
			filePositionFile, err = os.Create("filePosition")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// записываем нулевое значение в файл
			if _, err = filePositionFile.WriteString(string(os.SEEK_SET)); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	defer filePositionFile.Close()

	// открываем файл для чтения записей логов
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// читаем файловую позицию 
	var filePosition int 
	for scanner.Scan() {
		line := scanner.Text()
        fmt.Println(line)
        if line == "EOF" {
            break
        }
		// сохраняем текстовую ФП в int
		if filePosition, err := strconv.Atoi(line); err != nil {
			fmt.Println(err)
            os.Exit(1)
		}
	}

	// устанавливаем файловую позицию для чтения файла
	if _, err = file.Seek(int64(filePosition, os.SEEK_SET))
	if err!= nil {
        fmt.Println(err)
        os.Exit(1)
    }

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
