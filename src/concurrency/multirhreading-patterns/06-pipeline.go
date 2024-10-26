package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// generator creates channel with data
func generator(done chan struct{}, numbers *[]int) chan int {
	output := make(chan int)

	go func() {
		defer close(output)

		for _, num := range *numbers {
			select {
			case <-done:
				fmt.Println("generator(): done has been received")
				return
			case output <- num:
			}
		}
	}()

	return output
}

func add(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)

		for num := range inputCh {
			// Imitation of more expensive work
			delay := rand.Intn(1000)
			time.Sleep(time.Millisecond * time.Duration(delay))
			result := num + delay // some useful adding

			select {
			case <-doneCh:
				return
			case resultStream <- result:
			}
		}
	}()

	return resultStream
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)

		for num := range inputCh {
			// Imitation of more complex work
			delay := rand.Intn(1000)
			time.Sleep(time.Millisecond * time.Duration(delay))
			result := num * delay // some useful multiplying

			select {
			case <-doneCh:
				return
			case resultStream <- result:
			}
		}
	}()

	return resultStream
}

// fanOut() creates several goroutines for parallel data handling
func fanOut(doneCh chan struct{}, inputCh chan int, workers int) []chan int {
	resultChannels := make([]chan int, workers)

	for i := 0; i < workers; i++ {
		resultChannels[i] = add(doneCh, inputCh) // add() runs separate goroutine every iteration
	}

	return resultChannels
}

// fanIn() — combines the results of several channels into one
func fanIn(doneCh chan struct{}, channels ...chan int) chan int {
	finalStream := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		chCopy := ch
		wg.Add(1)

		go func() {
			defer wg.Done()

			for value := range chCopy {
				select {
				case <-doneCh:
					return
				case finalStream <- value:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(finalStream)
	}()

	return finalStream
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(doneCh, &numbers)

	// 'channels' = []chan int - slice of channels of int
	channels := fanOut(doneCh, inputCh, 10) // creates 10 goroutines with fanOut

	addResultCh := fanIn(doneCh, channels...) // combine results from all channels

	resultCh := multiply(doneCh, addResultCh) // pass results at next stage 'multiply'

	// output results
	for value := range resultCh {
		fmt.Printf("%d ", value)
	}

	fmt.Println()
}
