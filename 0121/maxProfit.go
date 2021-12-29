package _0121

func maxProfit(prices []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    res := 0
    n := len(prices)
    minPrice := prices[0]
    for i := 1; i < n; i++ {
        price := prices[i]
        res = max(res, price-minPrice)
        minPrice = min(minPrice, price)
    }

    return res
}
