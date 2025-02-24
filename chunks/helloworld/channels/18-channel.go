// https://habr.com/ru/post/490336/
// Работа с несколькими горутинами

package main

import (
	"fmt"
)

func square(ch chan int) {
	fmt.Println("[square] reading")
	num := <-ch
	ch <- num * num
}

func cube(ch chan int) {
	fmt.Println("[cube] reading")
	num := <-ch
	ch <- num * num * num
}

func main() {
	fmt.Println("[main] main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3

	fmt.Println("[main] sent testNum to squareChan")

	squareChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent testNum to cubeChan")

	cubeChan <- testNum

	fmt.Println("[main] resuming")
	fmt.Println("[main] from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of square and cube of", testNum, "is", sum)
	fmt.Println("[main] main() stopped")
}
