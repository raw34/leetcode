package _0377

func combinationSum4(nums []int, target int) int {
    res := 0
    dp := make([][]int, target+1)
    for i := 0; i < target+1; i++ {
        dp[i] = make([]int, target+1)
    }
    dp[0][0] = 1
    for i := 1; i < target+1; i++ {
        for j := 0; j < target+1; j++ {
            for _, num := range nums {
                if j >= num {
                    dp[i][j] += dp[i-1][j-num]
                }
            }
        }
        res += dp[i][target]
    }

    return res
}

func combinationSum42(nums []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1
    for i := 1; i < target+1; i++ {
        for _, num := range nums {
            if i >= num {
                dp[i] += dp[i-num]
            }
        }
    }

    return dp[target]
}
