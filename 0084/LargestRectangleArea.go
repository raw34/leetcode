package _084

import "math"

func largestRectangleArea(heights []int) int {
	area := float64(0)
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

			area = math.Max(area, float64(currW * currH))
		}

		stack = append(stack, i)
	}

	return int(area)
}