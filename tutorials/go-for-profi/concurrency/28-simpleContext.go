package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(time.Millisecond * 400)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Millisecond * 100):
		fmt.Println("f1():", r)
	}

	return
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Millisecond*100)
	defer cancel()

	go func() {
		time.Sleep(time.Millisecond * 400)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Millisecond * 100):
		fmt.Println("f2():", r)
	}

	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Millisecond * 100)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(time.Millisecond * 400)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Millisecond * 100):
		fmt.Println("f3():", r)
	}

	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <milliseconds>")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid delay:", os.Args[1])
		return
	}
	fmt.Println("Delay:", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}

// $ go run 28-simpleContext.go 4
