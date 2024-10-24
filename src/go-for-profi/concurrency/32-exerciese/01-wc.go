// Try to implement a competitive version of wc(1) that would use a buffered channel
// man (1) wc: wc - print newline, word, and byte counts for each file
// Steps for realize the task:
// open file
// create a buffered channel with a capacity of 10000
// read file line by line
// split line into words
// for each word, send it to the channel
// close the channel when the file is finished
// create a goroutine to read from the channel and count words and bytes
// create another goroutine to print the results
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func addLineToSls(linesSls *[]string, linesChan chan string) {
	for line := range linesChan {
		*linesSls = append(*linesSls, line)
	}
}

func main() {

	f := flag.String("f", "/etc/passwd", "File name")
	flag.Parse()
	fileName := *f
	fmt.Printf("Going to open and process file: %s\n", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

}
