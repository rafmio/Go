package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher()
	time.Sleep(time.Second * 2)
}

func sleepyGopher() {
	time.Sleep(time.Second * 1)
	fmt.Println(" ... snore ... ")
}
