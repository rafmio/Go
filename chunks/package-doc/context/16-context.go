// Generated by Preplexity
package main

import (
	"context"
	"fmt"
	"time"
)

// simulateWork simulates a time-consuming task.
func simulateWork(ctx context.Context, id int) {
	select {
	case <-time.After(2 * time.Second): // Simulate work taking 2 seconds
		fmt.Printf("Goroutine %d completed work\n", id)
	case <-ctx.Done(): // Handle cancellation
		fmt.Printf("Goroutine %d cancelled: %v\n", id, ctx.Err())
	}
}

func main() {
	// Using WithCancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure resources are cleaned up

	go simulateWork(ctx, 1)

	time.Sleep(1 * time.Second) // Let it run for a bit
	cancel()                    // Cancel the context
	time.Sleep(1 * time.Second) // Wait to see the cancellation message

	// Using WithTimeout
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go simulateWork(ctx, 2)

	time.Sleep(2 * time.Second) // Wait longer than timeout to see cancellation
	fmt.Println("Finished WithTimeout example")

	// Using WithDeadline
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go simulateWork(ctx, 3)

	time.Sleep(4 * time.Second) // Wait longer than deadline to see cancellation
	fmt.Println("Finished WithDeadline example")
}
