package _1863

func subsetXORSum(nums []int) int {
    length := len(nums)
    res := 0
    var backtrack func(path, begin int)
    backtrack = func(path, begin int) {
        res += path
        for i := begin; i < length; i++ {
            path ^= nums[i]
            backtrack(path, i+1)
            path ^= nums[i]
        }
    }
    backtrack(0, 0)

    return res
}
