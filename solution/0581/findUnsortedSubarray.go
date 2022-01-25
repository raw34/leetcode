package _0581

import "math"

func findUnsortedSubarray(nums []int) int {
    n := len(nums)
    minNum := math.MaxInt32
    maxNum := math.MinInt32
    left := -1
    right := -1
    for i := 0; i < n; i++ {
        if nums[i] >= maxNum {
            maxNum = nums[i]
        } else {
            right = i
        }
        if nums[n-i-1] <= minNum {
            minNum = nums[n-i-1]
        } else {
            left = n - i - 1
        }
    }

    res := right - left + 1
    if right == -1 {
        res = 0
    }

    return res
}

func findUnsortedSubarray2(nums []int) int {
    n := len(nums)
    stack := make([]int, 0)
    left := n
    for i := 0; i < n; i++ {
        curr := nums[i]
        for len(stack) > 0 && nums[stack[len(stack)-1]] > curr {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if top < left {
                left = top
            }
        }
        stack = append(stack, i)
    }
    right := 0
    for i := n - 1; i >= 0; i-- {
        curr := nums[i]
        for len(stack) > 0 && nums[stack[len(stack)-1]] < curr {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if top > right {
                right = top
            }
        }
        stack = append(stack, i)
    }

    res := right - left + 1
    if res < 0 {
        res = 0
    }

    return res
}
