// Try to implement a competitive version of wc(1) that would use shared memory
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type sharedMemory struct {
	linesCount   int
	wordsCount   int
	symbolsCount int
	mu           sync.Mutex
}

func (s *sharedMemory) incrementLines() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.linesCount++
}

func (s *sharedMemory) incrementWords(words int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.wordsCount += words
}

func (s *sharedMemory) incrementSymbols(symbols int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.symbolsCount += symbols
}

func readFileAndSendLineToChannel(file *os.File) (chan string, error) {
	linesChan := make(chan string)

	scanner := bufio.NewScanner(file)

	go func() error {
		defer close(linesChan)

		for scanner.Scan() {
			linesChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		return nil
	}()

	return linesChan, nil
}

func (s *sharedMemory) processLines(lines <-chan string) {
	var wg sync.WaitGroup

	for line := range lines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.incrementLines()
			s.incrementWords(len(strings.Fields(line)))
			s.incrementSymbols(len(line) + 1)
		}()
	}

	go func() {
		wg.Wait()
	}()
}

func (s *sharedMemory) getInfo() {
	fmt.Printf("Total lines: %d\n", s.linesCount)
	fmt.Printf("Total words: %d\n", s.wordsCount)
	fmt.Printf("Total symbols: %d\n", s.symbolsCount)
}

func main() {
	sharedMemory := &sharedMemory{}

	fileName := "/etc/passwd"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("ERROR opening file: %v\n", err)
		return
	}
	defer file.Close()

	linesChan, err := readFileAndSendLineToChannel(file)
	if err != nil {
		fmt.Printf("ERROR reading file: %v\n", err)
		return
	}

	sharedMemory.processLines(linesChan)
	sharedMemory.getInfo()
}
