package _0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)

    res := make([][]int, 0)

    var backtrack func(list []int, i int)
    backtrack = func(list []int, start int) {
        temp := make([]int, len(list))
        copy(temp, list)
        res = append(res, temp)

        for i := start; i < len(nums); i++ {
            if i != start && nums[i] == nums[i-1] {
                continue
            }
            list = append(list, nums[i])
            backtrack(list, i+1)
            list = list[0 : len(list)-1]
        }
    }
    backtrack([]int{}, 0)

    return res
}
