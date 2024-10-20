package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting main()...")

	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout c1")
	}

}
