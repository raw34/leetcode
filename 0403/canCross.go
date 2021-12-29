package _0403

func canCross(stones []int) bool {
    n := len(stones)
    // dp[i][k]表示第i个石头在上一次跳跃k长度时是否可达
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
    }
    // 初始化起点石头
    dp[0][0] = true
    // 从第二个石头开始遍历
    for i := 1; i < n; i++ {
        // 从第一个石头开始向当前石头遍历
        for j := 0; j < i; j++ {
            // 步长 = 当前石头和前一个石头的距离
            k := stones[i] - stones[j]
            // ???
            if k > j+1 {
                continue
            }
            dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
            // 当遍历到最后一个石头，且当该石头可达时返回
            if i == n-1 && dp[i][k] {
                return true
            }
        }
    }

    return false
}
