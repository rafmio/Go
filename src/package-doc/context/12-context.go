// https://marketsplash.com/tutorials/go/golang-context/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Context timed out:", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed")
	}
}
