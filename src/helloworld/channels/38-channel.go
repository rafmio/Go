// https://developers.sber.ru/link/gcsVFwahja
package main

import (
	"fmt"
	"time"
)

func longOperation() {
	time.Sleep(time.Second)
}

func main() {
	numbers := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numbers <- i
        }
        close(numbers
		}
	}
}
