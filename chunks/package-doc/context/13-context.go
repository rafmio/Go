// https://marketsplash.com/tutorials/go/golang-context/
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345)

	userID := ctx.Value("userID")

	if userID != nil {
		fmt.Println("User ID:", userID)
	} else {
		fmt.Println("User ID not found")
	}
}
