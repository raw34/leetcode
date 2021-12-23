package _1863

func subsetXORSum(nums []int) int {
    length := len(nums)
    res := 0
    path := 0
    var backtrack func(begin int)
    backtrack = func(begin int) {
        res += path
        for i := begin; i < length; i++ {
            path ^= nums[i]
            backtrack(i + 1)
            path ^= nums[i]
        }
    }
    backtrack(0)

    return res
}
