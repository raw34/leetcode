package _084

import (
	"github.com/raw34/leetcode/runtime"
)

func largestRectangleArea(heights []int) int {
	area := 0
	stack := make([]int, 0)
	newHeights := append([]int{0}, append(heights, 0)...)
	for i := 0; i < len(newHeights); i++ {
		for len(stack) > 0 && newHeights[i] < newHeights[stack[len(stack)-1]]{
			curr := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			currH := newHeights[curr]

			left := stack[len(stack)-1]
			right := i
			currW := right - left - 1

			area = runtime.Max(area, currW * currH)
		}

		stack = append(stack, i)
	}

	return area
}