package main

import (
	"fmt"
	"sync"
)

func main() {
	servernames := make([]string, 3)
	servernames[0] = "server1"
	servernames[1] = "server2"
	servernames[2] = "server3"

	stringNumbers := make([]string, len(servernames))
	stringNumbers[0] = "one"
	stringNumbers[1] = "two"
	stringNumbers[2] = "three"

	numsAndServersMap := make(map[string]map[string]int)

	var wg sync.WaitGroup
	var mu sync.Mutex

	// range servernames in separate goroutines
	for _, serverName := range servernames {

		wg.Add(1)
		// defer wg.Done()

		go func(serverName string) {
			mu.Lock()
			numsAndServersMap[serverName] = make(map[string]int)
			mu.Unlock()

			for innerIdx, num := range stringNumbers {

				wg.Add(1)
				// defer wg.Done()

				go func(serverName string, innderIdx string) {
					mu.Lock()
					numsAndServersMap[serverName][num] = innerIdx + 1
					mu.Unlock()

					wg.Done()
				}(serverName, num)

			}

			wg.Done()
		}(serverName)

	}

	wg.Wait()

	fmt.Println("Print result:")
	for serverName, nums := range numsAndServersMap {
		fmt.Println(serverName, nums)
	}

}
