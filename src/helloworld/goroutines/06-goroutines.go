package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)

		time.Sleep(time.Second * 1)
	}
}

func sleepyGopher(num int) {
	time.Sleep(time.Second * 1)
	fmt.Println(num, " ... snore ... ")
}
