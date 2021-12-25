package _0053

func maxSubArray(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    for i := 1; i < n; i++ {
        dp[i] = max(nums[i], dp[i-1]+nums[i])
    }

    res := nums[0]
    for i := 0; i < n; i++ {
        res = max(res, dp[i])
    }

    return res
}

func maxSubArray2(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(nums)
    res := nums[0]
    // dp[i]表示到索引i为止的最大子序列和
    dp := make([]int, n)
    dp[0] = nums[0]
    for i := 1; i < n; i++ {
        dp[i] = max(nums[i], dp[i-1]+nums[i])
        res = max(res, dp[i])
    }

    return res
}
