package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("The time is: ", currentTime)
	fmt.Println("currentTime.Unix(): ", currentTime.Unix())

	fmt.Println("year: ", currentTime.Year())
	fmt.Println("month: ", currentTime.Month())
	fmt.Println("hour: ", currentTime.Hour())
	fmt.Println("minute: ", currentTime.Minute())
	fmt.Println("second: ", currentTime.Second())
	fmt.Println("------------------------------")

	theTime := time.Date(2023, 3, 24, 9, 30, 45, 100, time.Local)
	fmt.Println("The time is: ", theTime)

	fmt.Println(theTime.Format("2006-1-2 15:4:5"))
	fmt.Println(theTime.Format("2006-01-02 03:04:05 pm"))
}
