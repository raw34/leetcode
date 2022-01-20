package _0401

import "fmt"

func readBinaryWatch(turnedOn int) []string {
    res := make([]string, 0)
    times := []int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
    var dfs func(h, m, begin, count int)
    dfs = func(h, m, begin, count int) {
        if h > 11 || m > 59 || count > turnedOn {
            return
        }

        if count == turnedOn {
            time := fmt.Sprintf("%d:%02d", h, m)
            res = append(res, time)
            return
        }

        for i := begin; i < 10; i++ {
            if h > 11 || m > 59 {
                continue
            }
            t := times[i]
            count++
            if i < 4 {
                h += t
                dfs(h, m, i+1, count)
                h -= t
            } else {
                m += t
                dfs(h, m, i+1, count)
                m -= t
            }
            count--
        }
    }
    dfs(0, 0, 0, 0)

    return res
}
