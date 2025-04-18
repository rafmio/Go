package main

import (
	"fmt"
	"time"
)

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, "...snore...")
}

func sleepyGopher2(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("sleepyGopher2", id, "...snore again")
}

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)

	for j := 0; j < 5; j++ {
		go sleepyGopher2(j)
	}
	time.Sleep(4 * time.Second)
}
