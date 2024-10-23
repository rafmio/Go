package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Background context cancelled")
			return
		default:
			fmt.Println("Background context not cancelled and running...")
		}
	}()

	deadLineCtx, deadLineCancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*200))
	defer deadLineCancel()

	go func() {
		select {
		case <-deadLineCtx.Done():
			fmt.Println("Context deadline exceeded")
		default:
			fmt.Println("Context not cancelled with deadline")
		}
	}()

	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer timeoutCancel()

	go func() {
		select {
		case <-timeoutCtx.Done():
			fmt.Println("Context timeout")
			return
		default:
			fmt.Println("Context was not cancelled with timeout")
		}
	}()

	time.Sleep(time.Second * 2)
}
