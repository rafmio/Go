package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	dateTime := time.Now()
	fmt.Println("dateTime:", dateTime)
	fmt.Println()

	fmt.Println(dateTime.Format("15:04"))
	fmt.Println(dateTime.Format("2006-1-2"))

	os.Exit(0)
}
