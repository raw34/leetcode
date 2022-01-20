package _0213

func rob(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }

    // dp[i]表示到当前家为止最大的收益
    dp := make([]int, n)

    // 第一种情况，偷第一家，不偷最后一家
    dp[0] = nums[0]
    dp[1] = nums[0]
    max1 := nums[0]
    for i := 2; i < n-1; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
        max1 = max(max1, dp[i])
    }

    // 第一种情况，不偷第一家，偷最后一家
    dp[0] = 0
    dp[1] = nums[1]
    max2 := nums[1]
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
        max2 = max(max2, dp[i])
    }

    return max(max1, max2)
}

func rob2(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }

    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    robMax := func(start, end int) int {
        dp := make([]int, 2)
        dp[0] = 0
        dp[1] = 0
        for i := start; i < end; i++ {
            dp[0], dp[1] = dp[1], max(dp[1], dp[0]+nums[i])
        }
        return dp[1]
    }

    // 第一种情况，偷第一家，不偷最后一家
    max1 := robMax(0, n-1)
    // 第一种情况，不偷第一家，偷最后一家
    max2 := robMax(1, n)

    return max(max1, max2)
}
