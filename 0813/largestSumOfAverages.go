package _0813

func largestSumOfAverages(nums []int, k int) float64 {
    max := func(a, b float64) float64 {
        if a < b {
            return b
        }
        return a
    }

    n := len(nums)
    // 预处理，计算前缀和
    sums := make([]int, n+1)
    for i := 1; i <= n; i++ {
        sums[i] = sums[i-1] + nums[i-1]
    }

    // dp[i][kk]表示到i位置的最大组合平均数
    dp := make([][]float64, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = make([]float64, k+1)
        dp[i][1] = float64(sums[i]) / float64(i)
        for kk := 2; kk <= k; kk++ {
            for j := 1; j < i; j++ {
                diff := float64(sums[i]-sums[j]) / float64(i-j)
                dp[i][kk] = max(dp[i][kk], dp[j][kk-1]+diff)
            }
        }
    }

    return dp[n][k]
}
