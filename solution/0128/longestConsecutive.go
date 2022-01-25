package _0128

func longestConsecutive(nums []int) int {
    max := func(a, b int) int {
        if a < b {
            return b
        }
        return a
    }
    hash := map[int]bool{}
    for _, num := range nums {
        hash[num] = true
    }
    res := 0
    for _, num := range nums {
        if !hash[num-1] {
            curr := num
            count := 1
            for hash[curr+1] {
                curr++
                count++
            }
            res = max(res, count)
        }
    }

    return res
}
