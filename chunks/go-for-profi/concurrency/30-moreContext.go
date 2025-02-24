package main

import (
	"context"
	"fmt"
)

type aKey string

// This function extracts the value from the context and checks if such a value exists
func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	myKey := aKey("mySecretValue")
	ctx := context.WithValue(context.Background(), myKey, "mySecretValue")

	searchKey(ctx, myKey)

	searchKey(ctx, aKey("notThere"))
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))
}
