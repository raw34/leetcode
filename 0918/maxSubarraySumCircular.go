package _0918

func maxSubarraySumCircular(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0]
    mx := nums[0]
    sum := nums[0]
    // 求最大子序列和
    for i := 1; i < n; i++ {
        dp[i] = nums[i] + max(dp[i-1], 0)
        mx = max(mx, dp[i])
        sum += nums[i]
    }

    mn := 0
    // 求nums[1]...nums[n-1]最小子序列和
    for i := 1; i < n-1; i++ {
        dp[i] = nums[i] + min(0, dp[i-1])
        mn = min(mn, dp[i])
    }

    return max(mx, sum-mn)
}
