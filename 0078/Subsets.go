package _0078

func subsets(nums []int) [][]int {
    res := make([][]int, 0)

    var backtrack func(list []int, i int)
    backtrack = func(list []int, i int) {
        temp := make([]int, len(list))
        copy(temp, list)
        res = append(res, temp)
        for j := i; j < len(nums); j++ {
            list = append(list, nums[j])
            backtrack(list, j+1)
            list = list[0 : len(list)-1]
        }
    }
    backtrack([]int{}, 0)

    return res
}
