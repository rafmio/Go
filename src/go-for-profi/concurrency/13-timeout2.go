package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(time.Second * 5)
		w.Wait()
	}()

	select {
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		println("Need a time duration!")
		return
	}

	var w sync.WaitGroup
	w.Add(1)

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		println("Invalid time duration!")
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	// duration := time.Duration(int32(t) * int32(time.Millisecond))
	fmt.Printf("Timeout period is %d\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timeout!")
	} else {
		fmt.Println("OK. Time elapsed.")
	}

	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timeout!")
	} else {
		fmt.Println("OK. Time elapsed.")
	}
}
