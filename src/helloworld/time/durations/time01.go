// https://www.geeksforgeeks.org/time-durations-in-golang/
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.IsZero())
	fmt.Println()

	var t2 time.Time
	fmt.Println(t2)
	fmt.Println(t2.IsZero())
}