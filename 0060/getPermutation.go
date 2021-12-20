package _0060

import "strconv"

func getPermutation(n int, k int) string {
    count := 0
    used := map[int]bool{}

    groupNum := 1
    for i := 1; i <= n; i++ {
        groupNum = groupNum * i
    }

    var backtrack func(path string) string
    backtrack = func(path string) string {
        if len(path) == n {
            count++
            if count == k {
                return path
            }
            return ""
        }

        groupNum = groupNum / (n - len(path))

        for i := 1; i <= n; i++ {
            if used[i] {
                continue
            }
            if k > groupNum {
                k = k - groupNum
                continue
            }

            path += strconv.Itoa(i)
            used[i] = true
            res := backtrack(path)
            path = path[:len(path)-1]
            used[i] = false
            if res != "" {
                return res
            }
        }
        return ""
    }
    return backtrack("")
}
