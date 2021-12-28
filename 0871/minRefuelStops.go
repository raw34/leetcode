package _0871

func minRefuelStops(target int, startFuel int, stations [][]int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    n := len(stations)
    // dp[i]表示加i次油可达公里数
    dp := make([]int, n+1)
    dp[0] = startFuel
    for i := 0; i < n; i++ {
        // 从后往前遍历加油站
        for j := i; j >= 0; j-- {
            // 当前加油次数可达该加油站时，计算下一次加油可达的公里数
            if dp[j] >= stations[i][0] {
                dp[j+1] = max(dp[j+1], dp[j]+stations[i][1])
            }
        }
    }

    // 遍历dp数组，发现最小符合条件次数返回
    for i := 0; i <= n; i++ {
        if dp[i] >= target {
            return i
        }
    }

    return -1
}
