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
	defer close(fileNamesChan)

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

	wg.Wait()

	fmt.Println(filesToRead, fileToWrite)

	bytesForWrite := make(chan []byte, len(filesToRead))
	defer close(bytesForWrite)

	for _, fileToRead := range filesToRead {
		wg.Add(1)
		go func(fileToRead string) {
			defer wg.Done()

			bytes, err := os.ReadFile(fileToRead)
			if err != nil {
				log.Println(err)
			}

			bytesForWrite <- bytes

			log.Println("bytes has been sent to channel")
		}(fileToRead)
	}

	fmt.Println("Before reading from channel...")

	go func() {
		for bytes := range bytesForWrite {
			err := os.WriteFile(fileToWrite, bytes, 0644)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("file has been written")
			}
		}
	}()
	fmt.Println("After reading from channel...")

	wg.Wait()
}
