package _0055

func canJump(nums []int) bool {
    n := len(nums)
    dp := make([]bool, n)
    dp[0] = true
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if dp[j] && j+nums[j] >= i {
                dp[i] = true
                break
            }
        }
    }

    return dp[n-1]
}
