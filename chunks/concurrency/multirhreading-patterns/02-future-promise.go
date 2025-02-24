// https://habr.com/ru/articles/852556/
// Receiving result and handling errors
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	value int
	err   error
}

func Promise(task func() (int, error)) chan Result {
	resultCh := make(chan Result, 1) // create channel for result

	go func() {
		value, err := task()                       // perform the task
		resultCh <- Result{value: value, err: err} // send result and error
		close(resultCh)                            // close the channel
	}()

	return resultCh // return channel to get result and error
}

func main() {
	// task that returns an error
	taskWithError := func() (int, error) {
		time.Sleep(1 * time.Second)
		num := rand.Intn(2)
		if num == 0 {
			return 42, nil
		} else {
			return 0, errors.New("Something went wrong")
		}
	}

	// run task via Promise
	future := Promise(taskWithError) // returns chan Result

	fmt.Println("Task is running, we can do something else...")

	// waiting for result and error
	result := <-future
	if result.err != nil {
		fmt.Printf("Task finished with error: %v\n", result.err)
	} else {
		fmt.Printf("Task finished, result: %d\n", result.value)
	}
}
