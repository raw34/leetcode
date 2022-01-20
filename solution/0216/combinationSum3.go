package _0216

func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0)
    var backtrack func(path []int, begin, sum int)
    backtrack = func(path []int, begin, sum int) {
        if len(path) == k && sum == n {
            temp := make([]int, k)
            copy(temp, path)
            res = append(res, temp)
        }

        if len(path) > k || sum > n {
            return
        }

        for i := begin; i <= 9; i++ {
            path = append(path, i)
            sum += i
            backtrack(path, i+1, sum)
            path = path[:len(path)-1]
            sum -= i
        }
    }
    backtrack([]int{}, 1, 0)

    return res
}
