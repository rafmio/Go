// Try to implement a competitive version of wc(1) that would use a buffered channel
// man (1) wc: wc - print newline, word, and byte counts for each file
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
<<<<<<< HEAD
	"time"
=======
>>>>>>> 6b6a28b (Home PC 24.10.2024 23:26 concurrency patterns)
)

type counts struct {
	object string // word or symbol
	count  int    // count of words or symbols
}

func openAndScanFile(fileName string, linesCounter *int) chan string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	linesChan := make(chan string)

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for scanner.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			*linesCounter++
			mu.Unlock()

			line := scanner.Text()

			linesChan <- line
		}()
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before return")

	return linesChan
}

func countWordsAnsSymbols(linesChan <-chan string) chan counts {
	countsChan := make(chan counts)
	var wg sync.WaitGroup

	for line := range linesChan {
		wg.Add(2)

		go func(line string) {
			defer wg.Done()
			symbolsCount := len(line) + 1 // 'wc' counts last '\n' symbol
			countsChan <- counts{object: "symbol", count: symbolsCount}
		}(line)

		go func(line string) {
			defer wg.Done()
			wordsCount := len(strings.Fields(line))
			countsChan <- counts{object: "word", count: wordsCount}
		}(line)
	}

	wg.Wait()

	return countsChan
}

func addWordsAndSymbolsVariables(wordsAnsSymbolCountsChan <-chan counts, wordsCounter, symbolsCounter *int) {
	for count := range wordsAnsSymbolCountsChan {
		switch count.object {
		case "word":
			*wordsCounter += count.count
		case "symbol":
			*symbolsCounter += count.count
		}
	}
}

func processWord(line string, wg *sync.WaitGroup, wordCountChan chan int, symbsCountChan chan int) {
	defer wg.Done()

	words := strings.Fields(line)
	wordCountChan <- len(words)

	symbsCountChan <- len(line) - len(strings.Replace(line, " ", "", -1))
}

func main() {

	f := flag.String("f", "/etc/passwd", "File name")
	flag.Parse()
	fileName := *f
	fmt.Printf("Going to open and process file: %s\n", fileName)

	var linesCounter int = 0
	var wordsCounter int = 0
	var symbolsCounter int = 0

<<<<<<< HEAD
	var wg sync.WaitGroup

	linesSls := make([]string, 0)
	linesChan := make(chan string)
	// defer close(linesChan)
	go addLineToSls(&linesSls, linesChan)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			linesChan <- scanner.Text()
		}()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Wait()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("len(linesSls)", len(linesSls))

	var linesCount int
	var wordsCount int
	var symbsCount int
	linesCountChan := make(chan int)
	wordsCountChan := make(chan int)
	symbsCountChan := make(chan int)

	for _, line := range linesSls {
		wg.Add(1)
		linesCountChan <- 1
		go processWord(line, &wg, wordsCountChan, symbsCountChan)
	}
	wg.Wait()

=======
	linesChan := openAndScanFile(fileName, &linesCounter)
	defer close(linesChan)

	wordsAnsSymbolCountsChan := countWordsAnsSymbols(linesChan)

	addWordsAndSymbolsVariables(wordsAnsSymbolCountsChan, &wordsCounter, &symbolsCounter)

	fmt.Printf("%d %d %d\n", linesCounter, wordsCounter, symbolsCounter)
>>>>>>> 6b6a28b (Home PC 24.10.2024 23:26 concurrency patterns)
}
