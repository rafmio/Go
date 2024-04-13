package main

import (
	"bufio"
	"fmt"
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
		if os.IsNotExist(err) {
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
			fmt.Println(err.Error())
			fmt.Printf("Type of err: %T\n", err)
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

	// сначала проверяем не превышает ли файловая позиция размера файла
	// если превышает - ФП = 0
	info, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileSize := info.Size()

	scanner := bufio.NewScanner(filePositionFile)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if line == "EOF" {
			break
		}
		// сохраняем текстовую ФП в int
		if filePosition, err = strconv.Atoi(line); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if filePosition > int(fileSize) {
		fmt.Println("filePosition: ", filePosition, "fileSize:", fileSize)
		fmt.Println("file position greater than file size")
		filePosition = 0
	}

	// debug:
	fmt.Println("current filePosition:", filePosition)

	// устанавливаем файловую позицию для чтения файла
	_, err = file.Seek(int64(filePosition), os.SEEK_SET)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// читаем файл начиная с заданной файловой позиции
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// находим файловую позицию после прочтения
	filePoisition, err := file.Seek(0, os.SEEK_CUR)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// усекаем (trucnate) файл для перезаписи
	filePositionFile, err = os.Create("filePosition")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = filePositionFile.WriteString(fmt.Sprint(filePoisition))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer filePositionFile.Close()
}
