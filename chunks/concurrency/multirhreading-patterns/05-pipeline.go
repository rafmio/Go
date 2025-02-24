// https://habr.com/ru/articles/852556/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			time.Sleep(time.Millisecond * 100)
			result := value + 2

			select {
			case <-doneCh: // if we need to end goroutine
				fmt.Println("add(): doneCh has received chancel signal")
				return
			case resultCh <- result: // send result
			}
		}
	}()

	return resultCh
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		for value := range inputCh {
			time.Sleep(time.Microsecond * 120)
			result := value * 3

			select {
			case <-doneCh:
				fmt.Println("multiply(): doneCh has received chancel signal")
				return
			case resultCh <- result:
			}
		}
	}()

	return resultCh
}

// generator sends data to channel
func generator(doneCh chan struct{}, numbers *[]int) chan int {
	outputCh := make(chan int)

	go func() {
		defer close(outputCh)

		for _, num := range *numbers {
			time.Sleep(time.Microsecond * 150)
			select {
			case <-doneCh:
				fmt.Println("generator(): doneCh has received chancel signal")
				return
			case outputCh <- num:
			}
		}
	}()

	return outputCh
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 300} // data we will handle

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(doneCh, &numbers) // run generator that sends numbers

	// pipeline stages
	addCh := add(doneCh, inputCh)
	resultCh := multiply(doneCh, addCh)

	// simulate cancellation
	go func() {
		signal := rand.Intn(2)
		delay := rand.Intn(160)
		if signal == 0 {
			return
		} else {
			time.Sleep(time.Millisecond * time.Duration(delay))
			doneCh <- struct{}{}
			fmt.Println("doneCh has been sent cancellation signal")
		}
	}()

	// print results
	for res := range resultCh {
		fmt.Printf("%d ", res)
	}

	fmt.Println()
}
