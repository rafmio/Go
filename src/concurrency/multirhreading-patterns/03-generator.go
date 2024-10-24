// https://habr.com/ru/articles/852556/
// Generator
/*
• Generator is a pattern for working with multithreading in Go.
• The producer-consumer scheme is used for parallel data processing.
• The generator generates data and sends it to the consumer through the channel.
• Sending and receiving data is blocked until both sides are ready.
• This approach minimizes delays and allows data to be processed as it arrives.
*/
package main

import "fmt"

func main() {
	// data will be send to channel
	items := []int{10, 20, 30, 40, 50}

	// receive channel with data from Generator
	dataChannel := generator(items)

	// consumer handles data from the channel
	process(dataChannel)
}

// Generator creates channel and run a goroutine for send data
func generator(items []int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch) // close channel after sending data

		// range data and send to channel
		for _, item := range items {
			ch <- item
		}
	}()

	return ch
}

// process receives the data from the channel and outputs it
func process(dataChannel chan int) {
	for item := range dataChannel {
		fmt.Printf("Received: %d\n", item)
	}
}
