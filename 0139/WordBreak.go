package _0139

func wordBreak(s string, wordDict []string) bool {
    n := len(s)
    dict := map[string]bool{}
    for _, word := range wordDict {
        dict[word] = true
    }

    dp := make(map[int]bool, 0)
    dp[0] = true
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            curr := s[j:i]
            if dp[j] && dict[curr] {
                dp[i] = true
            }
        }
    }

    return dp[n]
}
