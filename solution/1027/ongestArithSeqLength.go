package _1027

func longestArithSeqLength(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(nums)
    res := 0
    // dp定义：当前索引当前差值的情况的等差数列长度
    dp := make([]map[int]int, n)
    // 初始化0索引位置
    dp[0] = map[int]int{}
    for i := 1; i < n; i++ {
        dp[i] = map[int]int{}
        for j := 0; j < i; j++ {
            // 计算当前位置和后一个位置数字的差值
            diff := nums[i] - nums[j]
            // 当前位置和当前差值的数列不存在时初始化为1
            if _, ok := dp[j][diff]; !ok {
                dp[j][diff] = 1
            }
            // 更新后一个位置的数列长度
            dp[i][diff] = dp[j][diff] + 1
            // 更新全局最长数列长度
            res = max(res, dp[i][diff])
        }
    }

    return res
}
