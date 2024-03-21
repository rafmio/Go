// https://developers.sber.ru/link/gcsVFwahja
package main

import (
	"fmt"
	"math/rand"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	numbers := make(chan int, 2)
	defer close(numbers)

	go func() {
		numbers <- rand.Intn(5)
		numbers <- rand.Intn(6)
	}()

	go func() {
		numbers <- rand.Intn(4)
		numbers <- rand.Intn(3)
	}()

	for i := 0; i < 2; i++ {
		a := <-numbers
		b := <-numbers
		fmt.Printf("sum(%d, %d) = %d\n", a, b, sum(a, b))
	}
}
