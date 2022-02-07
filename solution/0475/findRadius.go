package _0475

import "sort"

func findRadius(houses []int, heaters []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    abs := func(x int) int {
        if x < 0 {
            return -x
        }
        return x
    }
    sort.Ints(houses)
    sort.Ints(heaters)
    res := 0
    m := len(houses)
    n := len(heaters)
    for i, j := 0, 0; i < m; i++ {
        curr := abs(houses[i] - heaters[j])
        for j < n-1 && abs(houses[i]-heaters[j]) >= abs(houses[i]-heaters[j+1]) {
            curr = min(curr, abs(houses[i]-heaters[j+1]))
            j++
        }
        res = max(res, curr)
    }

    return res
}
