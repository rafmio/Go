package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	fmt.Printf("type of nums: %T\n", nums)

	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 300, 500, 1000)

	nums := []int{10, 20, 30, 40, 50}
	sum(nums...)
}
