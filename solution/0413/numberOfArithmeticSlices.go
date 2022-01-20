package _0413

func numberOfArithmeticSlices(nums []int) int {
    res := 0
    n := len(nums)
    dp := make([]int, n)
    // 从第三个数字开始遍历，等差数列最少3个数字
    for i := 2; i < n; i++ {
        // 当前位置数字和前两个数字构成等差数列时，记录等差数列个数
        if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
            dp[i] = dp[i-1] + 1
        }
        // 累加等差数列个数
        res += dp[i]
    }

    return res
}
