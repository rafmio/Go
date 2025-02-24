package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a channel to receive results
	resultCh := make(chan int)

	// Start a goroutine that will run until the context is canceled
	go func() {
		result, err := doSomeWork(ctx)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		resultCh <- result
	}()

	// Wait for the result with a deadline
	ctx, cancel = context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	select {
	case result := <-resultCh:
		fmt.Println("Result:", result)
	case <-ctx.Done():
		fmt.Println("Timed out waiting for result:", ctx.Err())
	}
}

func doSomeWork(ctx context.Context) (int, error) {
	// Simulate some work
	time.Sleep(3 * time.Second)

	// Check if the context has been canceled
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
		return 42, nil
	}
}
