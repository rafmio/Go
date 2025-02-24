// https://www.geeksforgeeks.org/time-durations-in-golang/
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	fmt.Println(t.Clock())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println("----------------------")
	fmt.Println(t.Date())
	fmt.Println(t.Day())
	fmt.Println(t.Month())
	fmt.Println(t.Year())
	fmt.Println(t.YearDay())
	fmt.Println(t.Weekday())
	fmt.Println(t.ISOWeek())
	fmt.Println("----------------------")
	fmt.Println(t.String())
	fmt.Println(t.Unix())
	fmt.Println(t.Zone())
	fmt.Println(t.UnixNano())
}