package _0152

func maxProduct(nums []int) int {
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

    n := len(nums)
    dpMax := make([]int, n)
    dpMax[0] = nums[0]
    dpMin := make([]int, n)
    dpMin[0] = nums[0]
    for i := 1; i < n; i++ {
        mx := dpMax[i-1] * nums[i]
        mn := dpMin[i-1] * nums[i]
        dpMax[i] = max(mx, max(nums[i], mn))
        dpMin[i] = min(mn, min(nums[i], mx))
    }

    res := dpMax[0]
    for i := 1; i < len(dpMax); i++ {
        res = max(res, dpMax[i])
    }

    return res
}

func maxProduct2(nums []int) int {
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

    res := nums[0]
    n := len(nums)
    dpMax := nums[0]
    dpMin := nums[0]
    for i := 1; i < n; i++ {
        mx := dpMax * nums[i]
        mn := dpMin * nums[i]
        dpMax = max(mx, max(nums[i], mn))
        dpMin = min(mn, min(nums[i], mx))
        res = max(res, dpMax)
    }

    return res
}
