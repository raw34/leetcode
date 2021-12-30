package _0718

func findLength(nums1 []int, nums2 []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(nums1)
    n := len(nums2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }

    res := 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if nums1[i-1] == nums2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = 0
            }
            res = max(res, dp[i][j])
        }
    }

    return res
}

func findLength2(nums1 []int, nums2 []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    m := len(nums1)
    n := len(nums2)
    dp := make([]int, n+1)

    res := 0
    for i := 1; i <= m; i++ {
        for j := n; j > 0; j-- {
            if nums1[i-1] == nums2[j-1] {
                dp[j] = dp[j-1] + 1
            } else {
                dp[j] = 0
            }
            res = max(res, dp[j])
        }
    }

    return res
}
