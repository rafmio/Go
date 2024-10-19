// Создайте конвейер, который бы читал текстовые файлы, вычислял количество
// вхождений заданной фразы в каждом текстовом файле и подсчитывал общее
// количество вхождений этой фразы во всех файлах
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	patternFileName = "ufw.log"
)

func readDir(fileNamesChan chan<- string, readAndFillWg *sync.WaitGroup) {
	fmt.Println("readDir() starts...")
	readAndFillWg.Add(1)
	defer readAndFillWg.Done()

	dir, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var readDirWg sync.WaitGroup

	for _, file := range dir {
		// fmt.Println("file:", file.Name())
		readDirWg.Add(1)
		go func(fileName string) {
			if strings.Contains(fileName, patternFileName) && !file.IsDir() {
				fileNamesChan <- fileName
			}
			readDirWg.Done()
		}(file.Name())
	}

	// close(fileNamesChan)
	fmt.Println("DEBUG: from readDir()")
	readDirWg.Wait()
}

func fillFileNamesSls(fileNames *[]string, fileNamesChan <-chan string, readAndFillWg *sync.WaitGroup) {
	fmt.Println("DEBUG: from fillFileNamesSls()...")
	readAndFillWg.Add(1)
	defer readAndFillWg.Done()

	for fileName := range fileNamesChan {
		*fileNames = append(*fileNames, fileName)
	}
	fmt.Println("DEBUG: from fillFileNamesMap()")

}

func main() {
	fileNamesChan := make(chan string, 4)
	fileNames := make([]string, 0)

	var readAndFillWg sync.WaitGroup
	go readDir(fileNamesChan, &readAndFillWg)
	go fillFileNamesSls(&fileNames, fileNamesChan, &readAndFillWg)
	readAndFillWg.Wait()

	fmt.Printf("Type of fileNames: %T\n", fileNames)
	time.Sleep(time.Second * 1)
	fmt.Println("fileNames map:", fileNames)
}
