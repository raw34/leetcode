package _0413

func numberOfArithmeticSlices(nums []int) int {
    res := 0
    n := len(nums)
    dp := make([]int, n)
    for i := 2; i < n; i++ {
        if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
            dp[i] = dp[i-1] + 1
        }
        res += dp[i]
    }

    return res
}
