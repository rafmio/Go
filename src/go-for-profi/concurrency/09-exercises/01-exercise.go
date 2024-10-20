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
)

const (
	patternFileName = "ufw.log"
)

// func makeFileNamesChan(wg *sync.WaitGroup) <-chan string {
func makeFileNamesChan() <-chan string {
	fileNamesChan := make(chan string)

	fileNames, err := os.ReadDir(".")
	if err != nil {
		log.Println("ERROR ReadDir():", err.Error())
	}

	var wg sync.WaitGroup

	for _, file := range fileNames {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if file.IsDir() {
				return
			}

			if strings.Contains(file.Name(), patternFileName) {
				fileNamesChan <- file.Name()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(fileNamesChan)
	}()
	return fileNamesChan
}

func makeRawDataChan(fileNamesSls *[]string) <-chan []byte {
	rawDataChan := make(chan []byte)
	var wg sync.WaitGroup

	for _, fn := range *fileNamesSls {
		wg.Add(1)
		go func(fn string) {
			defer wg.Done()
			rawData, err := os.ReadFile(fn)
			if err != nil {
				log.Println("ERROR ReadFile():", err.Error())
				return
			}

			rawDataChan <- rawData
		}(fn)
	}

	go func() {
		wg.Wait()
		close(rawDataChan)
	}()

	return rawDataChan
}

func main() {
	fileNamesChan := makeFileNamesChan()
	fileNamesSls := make([]string, 0)

	counter := 7

	for fn := range fileNamesChan {
		fileNamesSls = append(fileNamesSls, fn)
		// fmt.Println(fn)
		counter--
	}

	if counter == 0 {
		fmt.Println("No more files to process. Len of the fileNamesSls:", len(fileNamesSls))
	}

	rawDataSls := make([][]byte, 0)
	rawDataChan := makeRawDataChan(&fileNamesSls)

	for rawData := range rawDataChan {
		rawDataSls = append(rawDataSls, rawData)
	}

	if len(rawDataSls) == 7 {
		fmt.Println("rawDataSls is complete. Len of the rawDataSls:", len(rawDataSls))
	} else {
		fmt.Println("rawDataSls is incomplete, len(rawDataSls):", len(rawDataSls))
	}
}
