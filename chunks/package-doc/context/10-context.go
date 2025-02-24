// https://marketsplash.com/tutorials/go/golang-context/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context canceled")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

}
