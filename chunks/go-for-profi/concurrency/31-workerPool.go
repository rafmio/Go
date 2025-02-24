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
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	makeWP(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}
