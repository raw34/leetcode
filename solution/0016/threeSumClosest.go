package _0016

import (
    "math"
    "sort"
)

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    abs := func(x int) int {
        if x < 0 {
            return -x
        }
        return x
    }
    n := len(nums)
    res := math.MaxInt32
    for i := 0; i < n; i++ {
        l, r := i+1, n-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == target {
                return sum
            }
            if abs(target-sum) < abs(target-res) {
                res = sum
            }
            if sum < target {
                l++
            } else if sum > target {
                r--
            }
        }
    }

    return res
}
