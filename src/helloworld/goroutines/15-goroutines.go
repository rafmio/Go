package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	patternToRead  = "text"
	patternToWrite = "result"
)

func main() {
	dir, err := os.ReadDir(".")
	if err != nil {
		log.Println(err)
	}

	filesToRead := make([]string, 0)
	var fileToWrite string
	fileNamesChan := make(chan string)

	var wg sync.WaitGroup

	for _, file := range dir {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if strings.Contains(file.Name(), patternToRead) {
				fileNamesChan <- file.Name()
				log.Println("has been sent to channel:", file.Name())
			}

			if strings.Contains(file.Name(), patternToWrite) {
				log.Println("*result*.txt:", file.Name())
				fileToWrite = file.Name()
			}
		}()
	}

	go func() {
		for fileToReadName := range fileNamesChan {
			filesToRead = append(filesToRead, fileToReadName)
		}
	}()

	go func() {
		wg.Wait()
		close(fileNamesChan)
	}()

	fmt.Println(filesToRead, fileToWrite)

}
