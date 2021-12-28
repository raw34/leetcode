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
    dp[0][0] = 0
    dp[0][1] = 1
    for i := 1; i < n; i++ {
        n1 := nums1[i]
        n2 := nums2[i]
        prevN1 := nums1[i-1]
        prevN2 := nums2[i-1]
        if n1 > prevN1 && n2 > prevN2 && n1 > prevN2 && n2 > prevN1 {
            dp[i][0] = min(dp[i-1][0], dp[i-1][1])
            dp[i][1] = dp[i][0] + 1
            continue
        }
        if n1 > prevN1 && n2 > prevN2 {
            dp[i][0] = dp[i-1][0]
            dp[i][1] = dp[i-1][1] + 1
        }
        if n1 > prevN2 && n2 > prevN1 {
            dp[i][0] = dp[i-1][1]
            dp[i][1] = dp[i-1][0] + 1
        }
    }

    return min(dp[n-1][0], dp[n-1][1])
}
