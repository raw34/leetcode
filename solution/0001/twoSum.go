package _0001

func twoSum(nums []int, target int) []int {
    n := len(nums)
    hash := map[int]int{}
    res := make([]int, 0)
    for i := 0; i < n; i++ {
        curr := nums[i]
        prev := target - curr
        if j, ok := hash[prev]; ok && i != j {
            res = []int{j, i}
            break
        }
        hash[curr] = i
    }

    return res
}
