package _0801

func minSwap(nums1 []int, nums2 []int) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    n := len(nums1)
    dp := make([][2]int, n)
    // dp[i][0]表示首位不换的情况下，当前位为止的最小交换次数
    dp[0][0] = 0
    // dp[i][0]表示首位交换的情况下，当前位为止的最小交换次数
    dp[0][1] = 1
    for i := 1; i < n; i++ {
        n1 := nums1[i]
        n2 := nums2[i]
        prevIndex := i - 1
        prevN1 := nums1[prevIndex]
        prevN2 := nums2[prevIndex]
        // 当两数组当前位都大于前一位，且交换后两数组仍能保持递增时
        if n1 > prevN1 && n2 > prevN2 && n1 > prevN2 && n2 > prevN1 {
            dp[i][0] = min(dp[prevIndex][0], dp[prevIndex][1])
            dp[i][1] = min(dp[prevIndex][0], dp[prevIndex][1]) + 1
            continue
        }
        // 当两数组当前位都大于前一位时，不换标记不便，交换标记累加
        if n1 > prevN1 && n2 > prevN2 {
            dp[i][0] = dp[prevIndex][0]
            dp[i][1] = dp[prevIndex][1] + 1
        }
        // 当数组1当前位大于数组2前一位且数组2当前位大于数组1前一位时，不换标记和交换标记对调，交换标记累加
        if n1 > prevN2 && n2 > prevN1 {
            dp[i][0] = dp[prevIndex][1]
            dp[i][1] = dp[prevIndex][0] + 1
        }
    }

    return min(dp[n-1][0], dp[n-1][1])
}
