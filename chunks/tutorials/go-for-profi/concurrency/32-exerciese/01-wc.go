// Try to implement a competitive version of wc(1) that would use a buffered channel
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type resInfo struct {
	linesCount   int
	wordsCount   int
	symbolsCount int
	linesSls     []string
}

type wordsAndSybmols struct {
	words   int
	symbols int
}

func readLine(file *os.File, linesCount *int) chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			*linesCount++
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Println("ERROR reading file:", err)
		}
	}()
	return ch
}

func (r *resInfo) addLinesToSls(lines chan string) {
	r.linesSls = make([]string, 0)

	for line := range lines {
		r.linesSls = append(r.linesSls, line)
	}
}

func (r *resInfo) processLinesSls() chan wordsAndSybmols {
	ch := make(chan wordsAndSybmols, 8)
	var wg sync.WaitGroup

	for _, line := range r.linesSls {
		wg.Add(1)
		go func() {
			ws := &wordsAndSybmols{
				symbols: len(line),
				words:   len(strings.Fields(line)),
			}
			ch <- *ws
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
		fmt.Printf("Finished processing lines\n")
	}()

	return ch
}

func (r *resInfo) incrementWordsAndSymbols(ch chan wordsAndSybmols) {
	for ws := range ch {
		r.wordsCount += ws.words
		r.symbolsCount += ws.symbols + 1
	}
}

func main() {
	r := new(resInfo)
	fileName := flag.String("f", "/etc/passwd", "File name")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Println("ERROR opening file:", err)
	} else {
		fmt.Printf("The  file %s has been opened\n", *fileName)
	}
	defer file.Close()

	lines := readLine(file, &r.linesCount)

	r.addLinesToSls(lines)
	results := r.processLinesSls() // chan wordsAndSybmols
	r.incrementWordsAndSymbols(results)

	fmt.Printf("Total lines: %d\n", r.linesCount)
	fmt.Printf("Total words: %d\n", r.wordsCount)
	fmt.Printf("Total symbols: %d\n", r.symbolsCount)
}
