// Create a pipeline to calculate the sum of the squares of all natural numbers
// in a given range
package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

func runRange(min, max int) <-chan int {
	if min < 0 || max < 0 {
		fmt.Println("ERROR: invalid min or max")
	}
	ch := make(chan int, max/2)
	var wg sync.WaitGroup

	for i := min; i <= max; i++ {
		wg.Add(1)
		go func(num int) {
			ch <- num
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch) // Close the channel when all numbers are processed
	}()

	return ch
}

func squaringNum(ch <-chan int) <-chan int {
	chSquared := make(chan int, len(ch))

	var wg sync.WaitGroup

	for num := range ch {
		wg.Add(1)
		go func(num int) {
			chSquared <- num * num
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		close(chSquared) // Close the channel when all squares are calculated
	}()

	return chSquared
}

func summingSquares(ch <-chan int) int {
	sum := 0
	for num := range ch {
		sum += num
	}
	return sum
}

func main() {
	min := flag.Int("min", 0, "Minimum value")
	max := flag.Int("max", 0, "Maximum value")

	flag.Parse()

	if *min < 0 || *max <= 0 {
		fmt.Println("Please provide both min and max values.")
		os.Exit(1)
	}

	ch := runRange(*min, *max)
	chSquared := squaringNum(ch)
	sum := summingSquares(chSquared)

	fmt.Printf("The sum of the squares of all natural numbers in the range [%d, %d] is: %d\n", *min, *max, sum)
}
