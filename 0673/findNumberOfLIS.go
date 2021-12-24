package _0673

func findNumberOfLIS(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    maxLen := 1
    n := len(nums)
    dp := make([]int, n)
    cnt := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = 1
        cnt[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                if dp[i] < dp[j]+1 {
                    dp[i] = dp[j] + 1
                    cnt[i] = cnt[j]
                } else if dp[i] == dp[j]+1 {
                    cnt[i] += cnt[j]
                }
            }
        }
        maxLen = max(maxLen, dp[i])
    }

    res := 0
    for i := 0; i < n; i++ {
        if dp[i] == maxLen {
            res += cnt[i]
        }
    }

    return res
}
