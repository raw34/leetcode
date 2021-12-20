package _0040

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    res := make([][]int, 0)

    sort.Ints(candidates)
    n := len(candidates)
    var backtrack func(path []int, sum, start int)
    backtrack = func(path []int, sum, start int) {
        if sum == target {
            temp := make([]int, len(path))
            copy(temp, path)
            res = append(res, temp)
        }

        for i := start; i < n; i++ {
            if sum > target {
                break
            }
            if i > start && candidates[i] == candidates[i-1] {
                continue
            }

            num := candidates[i]
            path = append(path, num)
            sum += num

            backtrack(path, sum, i+1)

            path = path[:len(path)-1]
            sum -= num
        }
    }
    backtrack([]int{}, 0, 0)

    return res
}
