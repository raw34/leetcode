package _0494

func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    diff := sum - target
    if diff < 0 || diff%2 != 0 {
        return 0
    }
    t := diff / 2

    //dp[i][j] 表示在数组 nums 的前 i 个数中选取元素，使得这些元素之和等于 j 的方案数。
    n := len(nums)
    dp := make([][]int, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]int, t+1)
    }
    dp[0][0] = 1
    for i := 1; i < n+1; i++ {
        num := nums[i-1]
        for j := 0; j < t+1; j++ {
            if j < num {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] + dp[i-1][j-num]
            }
        }
    }

    return dp[n][t]
}

func findTargetSumWays2(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    diff := sum - target
    if diff < 0 || diff%2 != 0 {
        return 0
    }
    t := diff / 2

    n := len(nums)
    dp := make([]int, t+1)
    dp[0] = 1
    for i := 1; i < n+1; i++ {
        num := nums[i-1]
        for j := t; j >= num; j-- {
            dp[j] += dp[j-num]
        }
    }

    return dp[t]
}
