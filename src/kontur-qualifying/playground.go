package main

import (
	"bufio"
	"fmt"
	"os"
)

type Playground struct {
	height int
	width  int
}

func (p Playground) Square() int {
	return p.height * p.width
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var width, height int
	fmt.Fscan(reader, &height)
	fmt.Fscan(reader, &width)

	rows := make([]string, 0)
	reader = bufio.NewReader(os.Stdin)
	for i := 0; i < height; i++ {
		var line string
		fmt.Fscan(reader, &line)
		rows = append(rows, line)
	}

	fmt.Println()
	for _, value := range rows {
		fmt.Println(value)
	}

	maxRec := maximalRectangle(rows)
	fmt.Println(maxRec)
}

// 85 largest rectangle
func maximalRectangle(matrix []string) int {
	var max int
	rows := len(matrix)
	if rows == 0 {
		return max
	}

	cols := len(matrix[0])
	nums := make([]int, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] != '0' {
				nums[col] = nums[col] + 1
			} else {
				nums[col] = 0
			}
		}

		maxArea := largestRectangleArea(nums)
		max = Max(max, maxArea)
	}

	return max
}

// 84 the largest rectangle in the histogram
func largestRectangleArea(heights []int) int {
	var max int
	var stack = make([]int, 0)
	var nums = make([]int, 0)

	nums = append(nums, []int{0}...)
	nums = append(nums, heights...)
	nums = append(nums, []int{0}...)

	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			tmp := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]

			max = Max(max, nums[tmp]*(i-stack[len(stack)-1]-1))
		}

		stack = append(stack, i)
	}

	return max
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
