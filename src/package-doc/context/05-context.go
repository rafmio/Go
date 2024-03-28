// https://www.golinuxcloud.com/golang-context/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	done := ctx.Done()

	for i := 0; ; i++ {
		select {
		case <-done:
			fmt.Println("done")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("tick", i)

		}
	}
}
