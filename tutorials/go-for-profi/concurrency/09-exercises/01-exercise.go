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

func makeEntriesChan(rawDataSls *[][]byte) <-chan string {
	entriesChan := make(chan string)
	var wg sync.WaitGroup

	for _, byteSls := range *rawDataSls {
		wg.Add(1)
		go func(byteSls []byte) {
			defer wg.Done()
			entries := strings.Split(string(byteSls), "\n")
			for _, entry := range entries {
				entriesChan <- entry
			}
		}(byteSls)
	}

	go func() {
		wg.Wait()
		close(entriesChan)
	}()

	return entriesChan
}

func makeSplittedStrChan(entriesSls *[]string) <-chan []string {
	splittedStrChan := make(chan []string, len(*entriesSls)/2)
	var wg sync.WaitGroup

	for _, entry := range *entriesSls {
		wg.Add(1)
		go func(entry string) {
			defer wg.Done()
			splittedStr := strings.Fields(entry)
			splittedStrChan <- splittedStr
		}(entry)
	}

	go func() {
		wg.Wait()
		close(splittedStrChan)
	}()

	return splittedStrChan
}

func makeIpPortTupleChan(splittedStrChan <-chan []string) <-chan []string {
	fmt.Println("Enter the makeIpProtTupleChan()")
	ipPortTupleChan := make(chan []string, len(splittedStrChan))

	var wg sync.WaitGroup

	for splittedStr := range splittedStrChan {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ipPortTuple := make([]string, 2)
			for _, filed := range splittedStr {
				if strings.HasPrefix(filed, "SRC=") {
					ipPortTuple[0] = strings.TrimPrefix(filed, "SRC=")
				}
				if strings.HasPrefix(filed, "DPT=") {
					ipPortTuple[1] = strings.TrimPrefix(filed, "DPT=")
				}
			}
			ipPortTupleChan <- ipPortTuple
		}()
	}

	go func() {
		wg.Wait()
		close(ipPortTupleChan)
	}()

	return ipPortTupleChan
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

	entriesSls := make([]string, 0)
	entriesChan := makeEntriesChan(&rawDataSls)

	for entry := range entriesChan {
		entriesSls = append(entriesSls, entry)
		// fmt.Println(entry)
	}

	fmt.Println("len(entriesSls):", len(entriesSls))

	uniqueIp := make(map[string]int)   // map[string(IP)]count
	uniquePort := make(map[string]int) // map[string(Port)]count

	splittedStrChan := makeSplittedStrChan(&entriesSls)

	ipPortTupleChan := makeIpPortTupleChan(splittedStrChan)

	counter = 0

	for ipPortTuple := range ipPortTupleChan {
		uniqueIp[ipPortTuple[0]]++
		uniquePort[ipPortTuple[1]]++
	}

	fmt.Println("len(uniqueIp):", len(uniqueIp))
	fmt.Println("len(uniquePort):", len(uniquePort))
}
