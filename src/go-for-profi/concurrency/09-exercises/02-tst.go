package main

import (
	"fmt"
	"sync"
)

func rangeAndSendToChan(wg *sync.WaitGroup, ch chan<- int) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- i
			wg.Done()
		}()
	}
	fmt.Println("End of fillSls()")
}

func fillSls(sls *[]int, ch <-chan int) {
	for num := range ch {
		*sls = append(*sls, num)
		// fmt.Printf("Received from channel: %d\n", num)
	}
	fmt.Println("End of fillSls()")
	// close(ch) // Close the channel after all goroutines are done sending.
	fmt.Println("End of main()")
	return // This is not necessary, but added for clarity.
	// The main() function will exit when the channel is closed.
	// The goroutine will continue to receive from the channel until it is closed.
	// However, it will not send anything to the channel after it is closed.
	// So, the range loop in the main() function will not receive anything.
	// This is a common pattern to ensure that goroutines are done before the main() function exits.
	// It's also a good practice to close the
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	sls := make([]int, 0)

	rangeAndSendToChan(&wg, ch)
	go fillSls(&sls, ch)

	wg.Wait()

	close(ch)

	fmt.Println(sls)
}
