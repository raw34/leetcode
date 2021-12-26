package _0198

func rob(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    res := nums[0]
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    for i := 1; i < n; i++ {
        if i == 1 {
            dp[i] = max(dp[i-1], nums[i])
        } else {
            dp[i] = max(dp[i-1], dp[i-2]+nums[i])
        }
        res = max(res, dp[i])
    }

    return res
}

func rob2(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(nums)
    prevMax := 0
    currMax := nums[0]
    for i := 1; i < n; i++ {
        prevMax, currMax = currMax, max(currMax, prevMax+nums[i])
    }

    return currMax
}
