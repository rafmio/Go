// P24 (*) Lotto: Draw N different random numbers from the set 1..M
package drawNnumbs

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDrawNnumbs(t *testing.T) {
	set := make([]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		set[i] = i + 1
	}

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(30)
	}

	for _, num := range nums {
		tstName := fmt.Sprintf("pass %d nums", num)
		t.Run(tstName, func(t *testing.T) {
			got := DrawNnumbs(set, num)
			if len(got) != num {
				t.Errorf("got %d, want %d", len(got), num)
			}
		})
	}
}
