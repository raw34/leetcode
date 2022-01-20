package _0046

func permute(nums []int) [][]int {
    length := len(nums)
    res := make([][]int, 0)
    visited := make([]bool, length)

    var backtrack func(list []int, visited []bool)
    backtrack = func(list []int, visited []bool) {
        if len(list) == length {
            temp := make([]int, length)
            copy(temp, list)
            res = append(res, temp)
            return
        }

        for i := 0; i < length; i++ {
            if visited[i] {
                continue
            }
            list = append(list, nums[i])
            visited[i] = true
            backtrack(list, visited)
            visited[i] = false
            list = list[0 : len(list)-1]
        }
    }
    backtrack([]int{}, visited)

    return res
}
