// Try to implement a competitive version of wc(1), which would use a control goroutine
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type fileInfo struct {
	fileName   string
	linesSls   []string
	linesCount int
	wordsCount int
	bytesCount int
	wordsChan  chan string
	bytesChan  chan string
}

func (f *fileInfo) fillLinesSls(inputCh <-chan string) {
	f.linesSls = make([]string, 0)
	for line := range inputCh {
		f.linesSls = append(f.linesSls, line)
	}
}

func (f *fileInfo) incrementLinesCount() {
	f.linesCount++
}

func (f *fileInfo) incrementWordsCount(line string) {
	words := len(strings.Fields(line))
	f.wordsCount += words
}

func (f *fileInfo) incrementBytesCount(line string) {
	f.bytesCount += len(line) + 1
}

func readFileAndSendLineToChannel(file *os.File) (chan string, error) {
	ch := make(chan string)
	var err error

	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			err = fmt.Errorf("error reading file: %v", err)
		}
	}()

	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (f *fileInfo) managingGoroutine(linesCh chan string) {
	f.wordsChan = make(chan string)
	f.bytesChan = make(chan string)

	go func() {
		for _, line := range f.linesSls {
			f.incrementLinesCount()
			f.wordsChan <- line
			f.bytesChan <- line
		}
	}()

	go func() {
		for {
			select {
			case <-linesCh:
				f.fillLinesSls(linesCh)
			case word := <-f.wordsChan:
				f.incrementWordsCount(word)
			case byteStr := <-f.bytesChan:
				f.incrementBytesCount(byteStr)
			default:
				close(f.bytesChan)
				close(f.wordsChan)
			}
		}
	}()

}

func main() {
	fileName := flag.String("f", "/etc/passwd", "File name")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	fi := new(fileInfo)

	linesChan, err := readFileAndSendLineToChannel(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fi.managingGoroutine(linesChan)

	<-time.After(time.Second * 3)
	fmt.Printf("lines: %d, words: %d, symbols: %d\n", fi.linesCount, fi.wordsCount, fi.bytesCount)
}
