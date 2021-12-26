package _0740

import "math"

func deleteAndEarn(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    n := len(nums)
    // 统计出数组中重复数字的和，并求出最大、最小值
    minNum := math.MaxInt32
    maxNum := math.MinInt32
    sums := make([]int, 0)
    for i := 0; i < n; i++ {
        num := nums[i]
        if len(sums) < num+1 {
            temp := make([]int, num+1)
            copy(temp, sums)
            sums = temp
        }
        minNum = min(minNum, num)
        maxNum = max(maxNum, num)
        sums[num] += num
    }

    // 利用打家劫舍思路求解
    prevMax := 0
    currMax := 0
    for i := minNum; i < maxNum+1; i++ {
        prevMax, currMax = currMax, max(currMax, prevMax+sums[i])
    }

    return currMax
}
