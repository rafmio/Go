/*
Change the Go code of the 31-workerPool.go in such a way as to save the results in a file.
When working with a file, use a mutex and a critical partition or a control routine that
will write data to disk
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	// for i := 0; i < n; i++ {
	// 	c := Client{i, i}
	// 	clients <- c
	// }

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c := Client{i, i}
			clients <- c
		}(i)
	}
	wg.Wait()

	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	// chech correct amount of argumets
	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	// convert arguments to integers and check if valid numbers
	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert arguments to integers and check if valid numbers
	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)
	finished := make(chan interface{})

	file, err := os.OpenFile("tst2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("file has been opened")
	}

	go func(file *os.File) {
		for d := range data {
			// fmt.Printf("Client ID: %d\tint: ", d.job.id)
			// fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
			_, err = file.WriteString(fmt.Sprintf("Client ID: %d int: %d square: %d\n", d.job.id, d.job.integer, d.square))
			if err != nil {
				fmt.Println("ERROR writing to file:", err)
				finished <- true
				return
			}
		}
		file.Close()
		finished <- true
	}(file)

	makeWP(nWorkers)
	fmt.Printf("finished: %v\n", <-finished)
}
