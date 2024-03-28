// https://marketsplash.com/tutorials/go/golang-context/
// WithDeadline
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(),
		time.Now().Add(2*time.Second))
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Context met deadline:", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("Operation timed out")
	}
}
