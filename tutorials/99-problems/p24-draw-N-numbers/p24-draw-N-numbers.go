// P24 (*) Lotto: Draw N different random numbers from the set 1..M
package drawNnumbs

import "math/rand"

// DrawNnumbs() method draws n different random numbers from the set 1..m
func DrawNnumbs(set []int, num int) []int {
	nums := make([]int, num)

	for i := 0; i < num; i++ {
		nums[i] = set[rand.Intn(1_000_000)]
	}

	return nums
}
