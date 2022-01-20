package _0135

func candy(ratings []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    n := len(ratings)
    // 初始化，每个孩子先分一个
    need := make([]int, n)
    for i := 0; i < n; i++ {
        need[i] = 1
    }
    // 从左至右检查糖果分配合理性
    for i := 1; i < n; i++ {
        if ratings[i] > ratings[i-1] {
            need[i] = need[i-1] + 1
        }
    }
    // 从右至左检查糖果分配合理性
    for i := n - 2; i >= 0; i-- {
        if ratings[i] > ratings[i+1] {
            need[i] = max(need[i], need[i+1]+1)
        }
    }

    // 统计糖果总数
    res := 0
    for i := 0; i < n; i++ {
        res += need[i]
    }

    return res
}
