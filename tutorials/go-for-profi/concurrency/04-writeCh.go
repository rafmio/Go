package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("inside writeToChannel(): first 'Println()':", x) // will be printed
	c <- x
	fmt.Println("inside writeToChannel(): second 'Println()':", x) // never will be printed
	close(c)
	fmt.Println(x) // never will be printed
}

func main() {
	c := make(chan int)

	go writeToChannel(c, 10)
	time.Sleep(time.Millisecond * 100)

	fmt.Println("inside main(): last 'Println()")
}
