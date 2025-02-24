package main

import (
	"errors"
	"log"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4}

	inputCh := generator(&input) // returns channel with data
	go consumer(inputCh)         // consumer that handles data

	time.Sleep(time.Millisecond * 100)
}

// sends data to channel and close it
func generator(input *[]int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range *input {
			inputCh <- data
		}
	}()

	return inputCh
}

// consumer receives data
func consumer(ch chan int) {
	for data := range ch {
		err := callDatabase(data)
		if err != nil {
			log.Println(err)
		}
	}
}

func callDatabase(data int) error {
	return errors.New("Database query error")
}
