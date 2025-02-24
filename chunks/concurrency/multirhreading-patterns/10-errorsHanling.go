package main

import (
	"errors"
	"log"
	"math/rand"
)

// Result - structure for storing results and errors
type Result struct {
	data int
	err  error
}

func main() {
	input := []int{1, 2, 3, 4}

	resultCh := make(chan Result)

	// launching consumer that will send results and errors
	go consumer(generator(input), resultCh)

	// reading results
	for res := range resultCh {
		if res.err != nil {
			log.Println("ERROR:", res.err)
		} else {
			log.Printf("Result: %d\n", res.data)
		}
	}
}

// generator - generates results for input data
func generator(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			inputCh <- data * 10
		}
	}()

	return inputCh
}

func consumer(inputCh chan int, resultCh chan Result) {
	defer close(resultCh)

	for data := range inputCh {
		resp, err := callDatabase(data)
		resultCh <- Result{data: resp, err: err}
	}
}

func callDatabase(data int) (int, error) {
	randErr := rand.Intn(3) // диапазон от 0 до 2. Т.е. каждый 3й запрос - ошибка

	if randErr == 1 {
		return 0, errors.New("Database connection error")
	}

	// return data, errors.New("Database query error")
	return data, nil
}
