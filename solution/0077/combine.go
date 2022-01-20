package _0077

func combine(n int, k int) [][]int {
    res := make([][]int, 0)

    var backtrack func(path []int, begin int)
    backtrack = func(path []int, begin int) {
        if len(path) == k {
            temp := make([]int, k)
            copy(temp, path)
            res = append(res, temp)
            return
        }

        for i := begin; i <= n; i++ {
            path = append(path, i)
            backtrack(path, i+1)
            path = path[:len(path)-1]
        }
    }
    backtrack([]int{}, 1)

    return res
}
