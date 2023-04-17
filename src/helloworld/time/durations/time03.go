// https://www.geeksforgeeks.org/time-durations-in-golang/
package main

import (
	"fmt"
	"time"
)

func main() {
	var d time.Duration = 1_000_000_000
	fmt.Println(d.Hours())
	fmt.Println(d.Minutes())
	fmt.Println(d.Seconds())
	fmt.Println(d.Milliseconds())
	fmt.Println(d.Microseconds())
	fmt.Println(d.String())
}