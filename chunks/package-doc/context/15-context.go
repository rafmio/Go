package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) error {
	defer ctx.Done()

	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Printf("Task %d running...\n", i)
			time.Sleep(time.Millisecond * 200)
		}
	}

	return nil
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*1000))
	defer cancel()

	done := make(chan struct{})
	defer close(done)

	go func() {
		err := longRunningTask(ctx)
		if err != nil && !errors.Is(err, context.Canceled) {
			fmt.Printf("Error in long running task: %v\n", err)
		}
	}()

	select {
	case <-done:
		fmt.Println("Long running task completed")
	case <-ctx.Done():
		fmt.Println("Long running task cancelled")
	}

}
