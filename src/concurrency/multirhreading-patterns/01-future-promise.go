// https://habr.com/ru/articles/852556/

// Promise — обещание предоставить результат операции в будущем.
// Оно используется для запуска задачи и получения объекта, который будет хранить
// результат выполнения задачи (он же async).

// Future — это объект, который позволяет проверять готовность результата и извлекать его,
// когда он будет доступен (он же await)
package main

import (
	"fmt"
	"time"
)

func Promise(task func() int) chan int {
	resultCh := make(chan int, 1) // создаем канал для результата

	go func() {
		result := task()   // performs the task
		resultCh <- result // send result to chan
		close(resultCh)
	}()

	return resultCh // return channel to get result
}

func main() {
	longRunningTask := func() int {
		time.Sleep(2 * time.Second)
		return 42 // return the result
	}

	// run task via Promise
	future := Promise(longRunningTask)

	fmt.Println("Task is running, we can do something else...")

	// waiting for result
	result := <-future
	fmt.Printf("Task finished, result: %d\n", result)
}
