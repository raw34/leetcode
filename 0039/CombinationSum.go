package _0039

func combinationSum(candidates []int, target int) [][]int {
    length := len(candidates)
    res := make([][]int, 0)

    var backtrack func(list []int, start, sum int)
    backtrack = func(list []int, start, sum int) {
        if sum >= target {
            if sum == target {
                temp := make([]int, len(list))
                copy(temp, list)
                res = append(res, temp)
            }

            return
        }

        for i := start; i < length; i++ {
            list = append(list, candidates[i])
            backtrack(list, i, sum+candidates[i])
            list = list[0 : len(list)-1]
        }
    }
    backtrack([]int{}, 0, 0)

    return res
}
