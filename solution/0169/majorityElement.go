package _0169

func majorityElement(nums []int) int {
    n := len(nums)
    hash := map[int]int{}
    res := 0
    for i := 0; i < n; i++ {
        curr := nums[i]
        hash[curr]++
        if hash[curr] >= n/2 && hash[curr] > hash[res] {
            res = curr
        }
    }

    return res
}
