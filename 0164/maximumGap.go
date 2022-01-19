package _0164

func maximumGap(nums []int) int {
    n := len(nums)
    if n < 2 {
        return 0
    }

    maxNum := nums[0]
    for _, num := range nums {
        if num > maxNum {
            maxNum = num
        }
    }

    buf := make([]int, n)
    for bit := 1; maxNum/bit > 0; bit *= 10 {
        cnt := make([]int, 10)
        for i := 0; i < n; i++ {
            num := (nums[i] / bit) % 10
            cnt[num]++
        }

        for i := 1; i < 10; i++ {
            cnt[i] += cnt[i-1]
        }

        for i := n - 1; i >= 0; i-- {
            num := (nums[i] / bit) % 10
            buf[cnt[num]-1] = nums[i]
            cnt[num]--
        }
        copy(nums, buf)
    }

    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }

    res := 0
    for i := 1; i < n; i++ {
        res = max(res, nums[i]-nums[i-1])
    }

    return res
}
