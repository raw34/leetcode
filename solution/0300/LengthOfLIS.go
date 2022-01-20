package _0300

func lengthOfLIS(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }

        return a
    }

    // dp定义：dp[i] 的值代表 nums 以 nums[i] 结尾的最长子序列长度
    n := len(nums)
    dp := make([]int, n)
    // 设定初始值
    dp[0] = 1
    res := 1
    // 循环计算每个位置的长度
    for i := 1; i < n; i++ {
        // 初始化当前位置长度为1
        dp[i] = 1
        // 从0到i遍历，看情况重新计算最大值
        for j := 0; j < i; j++ {
            // 当前数字如果比它之前的数字大，需要重新计算
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        // 更新全局最大值
        res = max(res, dp[i])
    }

    return res
}
