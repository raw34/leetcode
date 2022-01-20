package _0673

func findNumberOfLIS(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    res := 0
    maxLen := 0
    n := len(nums)
    dp := make([]int, n)
    cnt := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = 1
        cnt[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] >= nums[i] {
                continue
            }
            if dp[i] == dp[j]+1 {
                cnt[i] += cnt[j]
            } else if dp[i] < dp[j]+1 {
                dp[i] = dp[j] + 1
                cnt[i] = cnt[j]
            }
        }
        if dp[i] == maxLen {
            res += cnt[i]
        } else if dp[i] > maxLen {
            maxLen = dp[i]
            res = cnt[i]
        }
        maxLen = max(maxLen, dp[i])
    }

    return res
}
