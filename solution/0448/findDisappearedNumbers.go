package _0448

func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    res := make([]int, 0)
    for i := 0; i < n; i++ {
        // 计算当前元素应在的位置
        j := (nums[i] - 1) % n
        // 在目标位置上累加n
        nums[j] += n
    }
    for i := 0; i < n; i++ {
        // 当前位置小于n时，说明当前位置的数字没有被累加过，也就是目标数字
        if nums[i] <= n {
            res = append(res, i+1)
        }
    }

    return res
}

func findDisappearedNumbers2(nums []int) []int {
    n := len(nums)
    res := make([]int, 0)
    hash := map[int]bool{nums[0]: true}
    for i := 1; i < n; i++ {
        num := nums[i]
        hash[num] = true
    }
    for i := 1; i < n+1; i++ {
        if !hash[i] {
            res = append(res, i)
        }
    }

    return res
}
