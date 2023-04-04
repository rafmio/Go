package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var num int
	daysStr := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &num)
	for i := 0; i < num; i++ {
		var day string
		fmt.Fscan(reader, &day)
		daysStr = append(daysStr, day)
	}

	days := make([]int, 0)
	for _, value := range daysStr {
		day, _ := strconv.Atoi(value)
		days = append(days, day)
	}

	sortedDays := make([]int, len(days))
	copy(sortedDays, days)
	sort.Ints(sortedDays)

	var max, min int
	min = sortedDays[0]
	max = sortedDays[len(sortedDays)-1]

	fmt.Println(max, min)

	contrastLtR := make(map[int]int)
	for idx, value := range days {
		if value == min {
			contrastLtR[min] = idx
			break
		}
	}

	for idx, value := range days {
		if value == max {
			contrastLtR[max] = idx
		}
	}

	contrastRtL := make(map[int]int)
	for idx, value := range days {
		if value == max {
			contrastRtL[max] = idx
			break
		}
	}

	for idx, value := range days {
		if value == min {
			contrastRtL[min] = idx
		}
	}

	slsOutput := make([]int, 0)

	difLtR := contrastLtR[max] - contrastLtR[min]
	difRtL := contrastRtL[max] - contrastRtL[min]

	if difLtR > difRtL {
		slsOutput = append(slsOutput, contrastLtR[max]+1)
		slsOutput = append(slsOutput, contrastLtR[min]+1)
	} else {
		slsOutput = append(slsOutput, contrastRtL[max]+1)
		slsOutput = append(slsOutput, contrastRtL[min]+1)
	}
	fmt.Println(contrastLtR)
	fmt.Println(contrastRtL)
	fmt.Println(slsOutput)
}
