package isprime

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	primeNums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	cases := make(map[int]bool)

	for i := 0; i <= 100; i++ {
		cases[i] = false
		for _, val := range primeNums {
			if i == val {
				cases[i] = true
			}
		}
	}
}
