// https://golangify.com/goroutines
// Размер буфера канала
package main

import "fmt"
import "time"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4

	defer time.Sleep(time.Second * 2)
	fmt.Println("main() stopped")
}
