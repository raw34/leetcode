package _0300

func lengthOfLIS(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }

        return a
    }

    n := len(nums)
    dp := make([]int, n)
    dp[0] = 1
    res := 1
    for i := 1; i < n; i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        res = max(res, dp[i])
    }

    return res
}
