package _0056

import "sort"

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    n := len(intervals)
    left, right := intervals[0][0], intervals[0][1]
    inRegion := func(i int) bool {
        if i >= left && i <= right {
            return true
        }
        return false
    }
    res := make([][]int, 0)
    for i := 1; i < n; i++ {
        l, r := intervals[i][0], intervals[i][1]
        if inRegion(l) && inRegion(r) {
            continue
        } else if inRegion(l) && !inRegion(r) {
            right = r
        } else if !inRegion(l) && inRegion(r) {
            left = l
        } else {
            res = append(res, []int{left, right})
            left = l
            right = r
        }
    }

    if left != intervals[0][0] || right != intervals[0][1] || len(res) == 0 {
        res = append(res, []int{left, right})
    }

    return res
}

func merge2(intervals [][]int) [][]int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    n := len(intervals)
    res := make([][]int, 0)
    for i := 0; i < n; i++ {
        m := len(res)
        l, r := intervals[i][0], intervals[i][1]
        if m == 0 || res[m-1][1] < l {
            res = append(res, []int{l, r})
        } else {
            res[m-1][1] = max(res[m-1][1], r)
        }
    }

    return res
}
