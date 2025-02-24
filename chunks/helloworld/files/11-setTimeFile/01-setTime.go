package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filename := "text.txt"

	createTime := time.Now().Add(time.Hour * -10)
	accessTime := createTime

	err := os.Chtimes(filename, createTime, accessTime)
	if err != nil {
		fmt.Println("changing time:", err)
	}
}
