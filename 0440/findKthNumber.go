package _0440

func findKthNumber(n int, k int) int {
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    getNodeCount := func(curr int) int {
        total := 0
        next := curr + 1
        for curr <= n {
            total += min(n+1, next) - curr
            curr *= 10
            next *= 10
        }

        return total
    }
    curr := 1
    k -= 1
    for k > 0 {
        count := getNodeCount(curr)
        if k >= count {
            k -= count
            curr++
        } else {
            k -= 1
            curr *= 10
        }
    }

    return curr
}
