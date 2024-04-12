package main

import (
	"fmt"
	"time"
)

func main() {
	input := "2024-08-13"
	layout := "2006-01-02"

	t, _ := time.Parse(layout, input)

	fmt.Println(t)
	fmt.Println(t.Format("02-Jan-2006"))
}
