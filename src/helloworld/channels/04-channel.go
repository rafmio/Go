package main

import (
	"fmt"
	"os"
)

func main() {
	mychan1 := make(chan string)

	go func() {
		mychan1 <- "Kissy-Missy"
		mychan1 <- "Huggu-Wuggy"
		mychan1 <- "Tosy-Bosy"
		mychan1 <- "Vinnie-Pooh"

		close(mychan1)
	}()

	for res := range mychan1 {
		fmt.Println(res)
	}

	fmt.Println("----------------------")

	mychan2 := make(chan string, 4)

	mychan2 <- "My friend of mystery"
	mychan2 <- "Sad but true"
	mychan2 <- "Nothing else metter"
	mychan2 <- "Master"

	// fmt.Println("Length of the channel is: ", len(mychan2))

	fmt.Println("-------------------------")

	mychan3 := make(chan string, 5)

	mychan3 <- "My friend of mystery"
	mychan3 <- "Sad but true"
	mychan3 <- "Nothing else metter"
	mychan3 <- "Master"

	fmt.Println("Length of the channel is: ", len(mychan3))
	fmt.Println("Capacity of the clannel is: ", cap(mychan3))

	os.Exit(0)

}

// Channel
// for loop in the channel
// https://www.geeksforgeeks.org/channel-in-golang/
