// https://habr.com/ru/articles/852556/

/*
The principle of operation of Semaphore is simple:
- The semaphore counter determines how many goroutines can simultaneously work with a resource.
- When goroutine wants to start working with a resource, it increments the counter.
- If the counter has reached the maximum value, the goroutine is blocked and waits until
  one of the working goroutines frees up space by reducing the counter.
*/

package main

import (
	"log"
	"sync"
)

// Semaphore — a structure for managing the number of parallel goroutines
type Semaphore struct {
	semaCh chan struct{}
}

// NewSemaphore creates a new semaphore with the given maximum capacity
func NewSemaphore(maxCapacity int) *Semaphore {
	return &Semaphore{semaCh: make(chan struct{}, maxCapacity)}
}

// Acquire acquires a semaphore
func (s *Semaphore) Acquire() {
	s.semaCh <- struct{}{}
}

// Release releases a semaphore
func (s *Semaphore) Release() {
	<-s.semaCh
}

func main() {
	var wg sync.WaitGroup

	// creating a semaphore that will allow only two goroutines to work at the same time
	maxCapacity := 2
	semaphore := NewSemaphore(maxCapacity)

	// launch 10 goroutines:
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(taskID int) {
			defer wg.Done()
			defer semaphore.Release() // free space and decrement counter

			semaphore.Acquire() // acquire a semaphore before working

			// simulate goroutine's work
			log.Printf("Working process has been launched with ID: %d", taskID)
		}(i)
	}

	wg.Wait()
}

/*
Что происходит в коде:
- Мы создаем семафор, который позволяет только двум горутинам одновременно выполнять свою работу.
- В цикле запускается 10 горутин. Каждая из них перед началом работы "захватывает"
  семафор с помощью метода Acquire(). Если все места заняты, горутина ждет.
- Когда горутина завершает выполнение, она вызывает Release(), освобождая место для следующей горутины.
*/
