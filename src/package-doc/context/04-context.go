package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("ctx.Err():", ctx.Err())
	fmt.Println("ctx.Done():", ctx.Done())
	fmt.Println("ctx.Value(\"key\"):", ctx.Value("key"))
	fmt.Println("ctx.Deadline():")
	fmt.Print(ctx.Deadline())

	fmt.Println()
}
