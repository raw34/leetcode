package _0047

import "sort"

func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    length := len(nums)
    res := make([][]int, 0)
    vis := make([]bool, length)

    var backtrack func(list []int, vis []bool)
    backtrack = func(list []int, vis []bool) {
        if len(list) == length {
            temp := make([]int, length)
            copy(temp, list)
            res = append(res, temp)
            return
        }

        for i := 0; i < length; i++ {
            if vis[i] {
                continue
            }
            if i != 0 && nums[i] == nums[i-1] && !vis[i-1] {
                continue
            }
            list = append(list, nums[i])
            vis[i] = true
            backtrack(list, vis)
            vis[i] = false
            list = list[0 : len(list)-1]
        }
    }
    backtrack([]int{}, vis)

    return res
}
