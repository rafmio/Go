// https://golang.howtos.io/a-complete-guide-to-go-s-context-package/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// fmt.Printf("Type of ctx: %T\n", ctx)
	ctx = context.WithValue(ctx, "Kissy", "Missy") // try to send 100 (int) instead of Missjy
	// fmt.Printf("Type of ctx: %T\n", ctx)

	go worker(ctx)

	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context) {
	value, ok := ctx.Value("Kissy").(string) // make assertion (my own test)

	fmt.Println(ok)
	fmt.Println(value)
}
