package _0873

func lenLongestFibSubseq(arr []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(arr)
    // 缓存数组索引，以便后续查询
    index := map[int]int{}
    for i := 0; i < n; i++ {
        index[arr[i]] = i
    }
    // dp定义：从索引i到j的符合条件子序列个数
    dp := make([][]int, n)
    res := 0
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := i + 1; j < n; j++ {
            // 任意两个数都可以作为子序列的初始值
            dp[i][j] = 2
            diff := arr[j] - arr[i]
            // 寻找当前两数的前序数字
            k, ok := index[diff]
            // 如果前序数字存在，且数字小于当前第一位数字
            if arr[k] < arr[i] && ok {
                // 更新当前序列长度
                dp[i][j] = dp[k][i] + 1
                // 更新全局最长子序列长度
                res = max(res, dp[i][j])
            }
        }
    }

    return res
}
