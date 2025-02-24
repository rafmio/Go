package main

import (
	"fmt"
	"os"
)

func main() {
	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T\n", mychannel)

	mychannel1 := make(chan int)
	fmt.Println("Value of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1 %T\n", mychannel1)

	os.Exit(0)
}
