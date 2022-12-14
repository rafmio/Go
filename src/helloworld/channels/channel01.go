package main

import "fmt"

func main() {
	var mychannel chan int
	fmt.Println("Value of the channel:  ", mychannel)
	fmt.Println("Type of the channel: %T", mychannel)

	mychannel1 := make(chan int)
	fmt.Println("Value of the channel1: ", mychannel1)
	fmt.Println("Type of the channel1:  ", mychannel1)

}

// Channel
// https://www.geeksforgeeks.org/channel-in-golang/
