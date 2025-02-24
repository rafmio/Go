package main

import (
	"fmt"
	"time"
)

func main() {
	tm := time.Now()
	fmt.Println(tm)

	layout := "02-01-2006"
	date := tm.Format(layout)
	fmt.Println(date)

	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	} else {
		fmt.Printf("Type of parsedTime: %T\n", parsedTime)
	}

	fmt.Println(fmt.Sprintf("The current time is: %s\n", &parsedTime))
	fmt.Println(fmt.Sprintf("The current time is: %s\n", parsedTime))
}
