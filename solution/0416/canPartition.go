package _0416

func canPartition(nums []int) bool {
    n := len(nums)
    maxNum := nums[0]
    sum := 0
    for _, num := range nums {
        sum += num
        if num > maxNum {
            maxNum = num
        }
    }
    t := sum / 2
    if sum%2 != 0 || t < maxNum {
        return false
    }

    // dp[i][j]表示对nums的前i个元素进行选取，dp记录是否存在选取值相加结果为目标值j
    dp := make([][]bool, n+1)
    for i := 0; i < n+1; i++ {
        dp[i] = make([]bool, t+1)
        dp[i][0] = true
    }
    for i := 1; i < n+1; i++ {
        num := nums[i-1]
        for j := 0; j < t+1; j++ {
            if j < num {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-num]
            }
        }
    }

    return dp[n][t]
}

func canPartition2(nums []int) bool {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    t := sum / 2
    if sum%2 != 0 {
        return false
    }

    dp := make([]bool, t+1)
    dp[0] = true
    for _, num := range nums {
        for i := t; i >= num; i-- {
            dp[i] = dp[i] || dp[i-num]
        }
    }

    return dp[t]
}
